from __future__ import annotations

import math
import operator

from dataclasses import dataclass, field
from enum import Enum, IntEnum, auto

from typing_extensions import Any, Callable

from ... import utils
from ..translator.translator import Translator
from ...types.serialization.noro6.fleet.fleet import Fleet
from ...types.serialization.noro6.fleet.fleet_info import FleetInfo

from ...logger import LOGGER


class _ORDER_EQUIPMENT_TRANSLATE(IntEnum):
    A = 0
    """Equipment 1"""
    B = 1
    """Equipment 2"""
    C = 2
    """Equipment 3"""
    D = 3
    """Equipment 4"""
    E = 4
    """Equipment 5"""
    X = 99
    """Equipment Extra Slots"""


_ORDER_EQUIPMENT_TRANSLATE_NAME_SET = {
    k for k in _ORDER_EQUIPMENT_TRANSLATE.__members__.keys()
}
_ORDER_EQUIPMENT_TRANSLATE_VALUE_SET = {
    str(v.value) for v in _ORDER_EQUIPMENT_TRANSLATE
}
_ORDER_EQUIPMENT_TRANSLATE_SET = {
    *_ORDER_EQUIPMENT_TRANSLATE_NAME_SET,
    *_ORDER_EQUIPMENT_TRANSLATE_VALUE_SET,
}


class _ORDER_SHIP_TRANSLATE(IntEnum):
    A = 0
    """Single Fleet / JTF Flag Ship"""
    B = 1
    """Single Fleet / JTF 2nd ship"""
    C = 2
    """Single Fleet / JTF 3rd ship"""
    D = 3
    """Single Fleet / JTF 4th ship"""
    E = 4
    """Single Fleet / JTF 5th ship"""
    F = 5
    """Single Fleet / JTF 6th ship"""
    G = 6
    """Single Fleet Vanguard 7th ship"""
    U = 6
    """JTF Escort Fleet Flag Ship"""
    V = 7
    """JTF Escort Fleet 2nd Ship"""
    W = 8
    """JTF Escort Fleet 3rd Ship"""
    X = 9
    """JTF Escort Fleet 4th Ship"""
    Y = 10
    """JTF Escort Fleet 5th Ship"""
    Z = 11
    """JTF Escort Fleet 6th Ship"""


_ORDER_SHIP_TRANSLATE_NAME_SET = {k for k in _ORDER_SHIP_TRANSLATE.__members__.keys()}
_ORDER_SHIP_TRANSLATE_VALUE_SET = {str(v.value) for v in _ORDER_SHIP_TRANSLATE}
_ORDER_SHIP_TRANSLATE_SET = {
    *_ORDER_SHIP_TRANSLATE_NAME_SET,
    *_ORDER_SHIP_TRANSLATE_VALUE_SET,
}


class MacroValueType(Enum):
    MNEMONIC = auto()
    ATTRIBUTE_ACCESS = auto()
    UNKNOWN = auto()


@dataclass(slots=True)
class Macro:
    value: str = ""
    type: MacroValueType = MacroValueType.UNKNOWN

    def eval(self) -> str:
        result: str = self.value

        if self.type is MacroValueType.MNEMONIC:
            ...
        elif self.type is MacroValueType.ATTRIBUTE_ACCESS:
            ...

        return result


def isValidMacro(val: str) -> Macro | None:
    if len(val) < 2:
        return None

    if val[0] != "<" and val[-1] != ">":
        return None

    macroStr = val.removeprefix("<").removesuffix(">")

    if all(v.isupper() or v == "_" for v in macroStr):
        # <MNEMONIC>
        return Macro(
            macroStr,
            (
                MacroValueType.MNEMONIC
                if all(len(v) for v in macroStr.split("_"))
                else MacroValueType.UNKNOWN
            ),
        )

    if len(macroStr) < 2:
        return None

    if macroStr[0] != "<" and macroStr[-1] != ">":
        return None

    macroStr = macroStr.removeprefix("<").removesuffix(">")

    if all(v in utils.ASCII_LETTER_SET or v == "." or v.isdecimal() for v in macroStr):
        # <<ATTRIBUTE.ACCESS>>
        return Macro(
            macroStr,
            (
                MacroValueType.ATTRIBUTE_ACCESS
                if all(len(v) for v in macroStr.split("_"))
                else MacroValueType.UNKNOWN
            ),
        )

    return Macro(val, MacroValueType.UNKNOWN)


def attrAccess(fleetInfo: FleetInfo, macro: Macro) -> str:
    if macro.type is not MacroValueType.ATTRIBUTE_ACCESS:
        return ""

    result: str = ""

    match macroSplit := macro.value.split("."):
        case ["Ship", posRaw, *attrs] if posRaw in _ORDER_SHIP_TRANSLATE_SET and len(
            attrs
        ):
            LOGGER.debug(f"get Ship[{posRaw = }. {attrs = }]")

            pos = (
                getattr(_ORDER_SHIP_TRANSLATE, posRaw)
                if posRaw in _ORDER_SHIP_TRANSLATE_NAME_SET
                else _ORDER_SHIP_TRANSLATE(int(posRaw))
            )

            targetShip = (
                fleetInfo.unionFleet.ships[pos]
                if fleetInfo.isUnion and fleetInfo.unionFleet is not None
                else fleetInfo.mainFleet().ships[pos]
            )

            match attrs:
                case [
                    "equipment",
                    equipmentPosRaw,
                    *equipmentAttrs,
                ] if equipmentPosRaw in _ORDER_EQUIPMENT_TRANSLATE_SET and len(
                    equipmentAttrs
                ) >= 1:
                    LOGGER.debug(f"{equipmentPosRaw = } {equipmentAttrs = }")
                    equipmentPos = (
                        getattr(_ORDER_EQUIPMENT_TRANSLATE, equipmentPosRaw)
                        if equipmentPosRaw in _ORDER_EQUIPMENT_TRANSLATE_NAME_SET
                        else _ORDER_EQUIPMENT_TRANSLATE(int(equipmentPosRaw))
                    )

                    _r = None

                    targetEquipment = (
                        targetShip.items[equipmentPos]
                        if equipmentPosRaw != "X"
                        or equipmentPos != _ORDER_EQUIPMENT_TRANSLATE.X
                        else targetShip.exItem
                    )
                    LOGGER.debug(f"{targetEquipment = }")

                    try:
                        _r = operator.attrgetter(".".join(equipmentAttrs))(
                            targetEquipment
                        )
                    except Exception as e:
                        LOGGER.error(f"{e}")
                        _r = None

                    if _r is not None:
                        result = str(_r) if not isinstance(_r, str) else _r
                        LOGGER.debug(f"{result = }")
                case _:
                    _r = None

                    try:
                        _r = operator.attrgetter(".".join(attrs))(targetShip)
                    except Exception as e:
                        LOGGER.error(f"{e = }")
                        _r = None

                    if _r is not None:
                        result = str(_r) if not isinstance(_r, str) else _r
                        LOGGER.debug(f"{result = }")

        case ["Fleet", fleet, *attrs] if fleet in {"A", "B", "U"} and len(attrs):
            if not fleetInfo.isUnion and fleet == "U":
                LOGGER.debug("Try to access union fleet but there isn't!")
                return ""

            targetFleet = {
                "A": fleetInfo.fleets[0],
                "B": fleetInfo.fleets[1],
                "C": fleetInfo.fleets[2],
                "D": fleetInfo.fleets[3],
                "U": fleetInfo.unionFleet,
            }.get(fleet)

            if targetFleet is None:
                LOGGER.error(f"Unknown {targetFleet = }")
                return ""

            match attrs:
                case ["los", level] if level in {"A", "B", "C", "D"}:
                    _lookUpSplit = {"A": 0, "B": 1, "C": 2, "D": 3}.get(level, None)

                    if _lookUpSplit is None:
                        LOGGER.error(f"unknown {level = }")
                        result = ""
                    else:
                        result = (
                            str(
                                Fleet.getUnionScoutScore(
                                    targetFleet, fleetInfo.admiralLevel
                                )[_lookUpSplit]
                            )
                            if fleetInfo.isUnion
                            else str(
                                Fleet.getScoutScore(
                                    targetFleet.ships, fleetInfo.admiralLevel
                                )[_lookUpSplit]
                            )
                        )
                        LOGGER.debug(f"{result = }")
                case _:
                    try:
                        _r = operator.attrgetter(".".join(attrs))(targetFleet)
                    except Exception as e:
                        LOGGER.error(e)
                        _r = None

                    if _r is not None:
                        result = str(_r) if not isinstance(_r, str) else _r
                        LOGGER.debug(f"{result = }")

        case _:
            LOGGER.debug(f"unknown pattern: {macroSplit!r}")

    return result


@dataclass(slots=True)
class PreDefineMacro:
    fleetInfo: FleetInfo
    translator: Translator

    _macroLookUpCache: dict[str, str] = field(init=False)
    _latexLookUpCache: dict[str, str] = field(init=False)

    def __post_init__(self):
        self._macroLookUpCache = {}
        self._latexLookUpCache = {}

        def _template(
            _pos: str,
            _latex: str,
            _macro: str,
            _access: str,
            _default: str = "",
            _functionWrapper: Callable[[Any], str] | None = None,
        ):
            _latex = _latex.format(_pos)
            _macro = _macro.format(_pos)
            _access = _access.format(_pos)
            try:
                _attrAccessResult = self._attrAccess(
                    Macro(_access, MacroValueType.ATTRIBUTE_ACCESS)
                )
                _accessResult = (
                    _attrAccessResult
                    if _functionWrapper is None
                    else _functionWrapper(_attrAccessResult)
                )
                LOGGER.debug(f"{_accessResult = }")
                self._macroLookUpCache.update({_macro: _accessResult})
                self._latexLookUpCache.update({_latex: _accessResult})
            except IndexError as e:
                LOGGER.debug(f"{e = } {_pos = }")
                self._macroLookUpCache.update({_macro: _default})
                self._latexLookUpCache.update({_latex: _default})

        for fleetPos in {"A", "B", "C", "D", "U"}:
            LOGGER.debug(f"{fleetPos = }")

            for losPos in {"A", "B", "C", "D"}:
                _template(
                    losPos,
                    rf"\fleet{fleetPos}los{{}}",
                    f"FLEET_{fleetPos}_LOS_{{}}",
                    f"Fleet.{fleetPos}.los.{{}}",
                    _functionWrapper=lambda v: str(
                        math.floor(100 * utils.convert(v, float, 0)) / 100
                    ),
                )
            _template(
                fleetPos,
                r"\fleet{}fullAirPower",
                "FLEET_{}_FULL_AIRPOWER",
                "Fleet.{}.fullAirPower",
            )

        for shipPos in _ORDER_SHIP_TRANSLATE_NAME_SET:
            LOGGER.debug(f"{shipPos = }")
            _template(shipPos, r"\ship{}nameJp", "SHIP_{}_NAME_JP", "Ship.{}.data.name")
            _template(
                shipPos,
                r"\ship{}nameEn",
                "SHIP_{}_NAME_EN",
                "Ship.{}.data.name",
                _functionWrapper=self.translator.translate_ship,
            )
            _template(shipPos, r"\ship{}level", "SHIP_{}_LEVEL", "Ship.{}.level")
            _template(
                shipPos,
                r"\ship{}fullAirPower",
                "SHIP_{}_FULL_AIRPOWER",
                "Ship.{}.fullAirPower",
            )

            for equipmentPos in _ORDER_EQUIPMENT_TRANSLATE_NAME_SET:
                LOGGER.debug(f"{equipmentPos = }")

                _template(
                    equipmentPos,
                    rf"\ship{shipPos}equipment{{}}nameJp",
                    f"SHIP_{shipPos}_EQUIPMENT_{{}}_NAME_JP",
                    f"Ship.{shipPos}.equipment.{{}}.data.name",
                )
                _template(
                    equipmentPos,
                    rf"\ship{shipPos}equipment{{}}nameEn",
                    f"SHIP_{shipPos}_EQUIPMENT_{{}}_NAME_EN",
                    f"Ship.{shipPos}.equipment.{{}}.data.name",
                    _functionWrapper=self.translator.translate_equipment,
                )
                _template(
                    equipmentPos,
                    rf"\ship{shipPos}equipment{{}}remodel",
                    f"SHIP_{shipPos}_EQUIPMENT_{{}}_REMODEL",
                    f"Ship.{shipPos}.equipment.{{}}.remodel",
                )
                _template(
                    equipmentPos,
                    rf"\ship{shipPos}equipment{{}}levelAlt",
                    f"SHIP_{shipPos}_EQUIPMENT_{{}}_LEVEL_ALT",
                    f"Ship.{shipPos}.equipment.{{}}.levelAlt",
                )

    @property
    def macroLookUp(self) -> dict[str, str]:
        return self._macroLookUpCache

    @property
    def latexLoopUp(self) -> dict[str, str]:
        return self._latexLookUpCache

    def _attrAccess(self, macro: Macro):
        return attrAccess(self.fleetInfo, macro)
