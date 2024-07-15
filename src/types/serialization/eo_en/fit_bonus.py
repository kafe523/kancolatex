from __future__ import annotations

from dataclasses import dataclass, field as std_field

from pydantic import BaseModel, Field
from typing_extensions import Optional

from ...equipment_id import EquipmentId
from ...const import EquipmentTypes, CountryType
from ...ship_id import ShipId


class FitBonusValue(BaseModel):
    firepower: int = Field(alias="houg", default=0)
    torpedo: int = Field(alias="raig", default=0)
    antiAir: int = Field(alias="tyku", default=0)
    armor: int = Field(alias="souk", default=0)
    evasion: int = Field(alias="kaih", default=0)
    asw: int = Field(alias="tais", default=0)
    los: int = Field(alias="saku", default=0)

    accuracy: int = Field(alias="houm", default=0)
    """Visible acc fit actually doesn't work according to some studies"""

    range: int = Field(alias="leng", default=0)
    bombing: int = Field(alias="baku", default=0)

    fromTypeId: int = Field(init=False, default=0)
    """Special field only use in noro6"""

    def __mul__(self, other: int) -> FitBonusValue:
        return FitBonusValue(
            houg=self.firepower * other,
            raig=self.torpedo * other,
            tyku=self.antiAir * other,
            souk=self.armor * other,
            kaih=self.evasion * other,
            tais=self.asw * other,
            saku=self.los * other,
            houm=self.accuracy * other,
            leng=self.range * other,
            baku=self.bombing * other,
        )

    def __add__(self, other: FitBonusValue) -> FitBonusValue:
        return FitBonusValue(
            houg=self.firepower + other.firepower,
            raig=self.torpedo + other.torpedo,
            tyku=self.antiAir + other.antiAir,
            souk=self.armor + other.armor,
            kaih=self.evasion + other.evasion,
            tais=self.asw + other.asw,
            saku=self.los + other.los,
            houm=self.accuracy + other.accuracy,
            leng=self.range + other.range,
            baku=self.bombing + other.bombing,
        )

    def hasBonus(self) -> bool:
        return any(
            (
                self.firepower,
                self.torpedo,
                self.antiAir,
                self.armor,
                self.evasion,
                self.asw,
                self.los,
                self.accuracy,
                self.range,
                self.bombing,
            )
        )

    def __eq__(self, value: object) -> bool:
        if type(value) is not type(self):
            return False

        return all(
            (
                value.firepower == self.firepower,
                value.torpedo == self.torpedo,
                value.antiAir == self.antiAir,
                value.armor == self.armor,
                value.evasion == self.evasion,
                value.asw == self.asw,
                value.los == self.los,
                value.accuracy == self.accuracy,
                value.range == self.range,
                value.bombing == self.bombing,
            )
        )

    def __hash__(self) -> int:
        return hash(
            (
                self.firepower,
                self.torpedo,
                self.antiAir,
                self.armor,
                self.evasion,
                self.asw,
                self.los,
                self.accuracy,
                self.range,
                self.bombing,
            )
        )


class FitBonusData(BaseModel):
    shipClass: Optional[list[int]] = Field(alias="shipClass", default=None)

    shipMasterIds: Optional[list[ShipId]] = Field(alias="shipX", default=None)
    """Master id = exact id of the ship"""

    shipIds: Optional[list[ShipId]] = Field(alias="shipS", default=None)
    """Base id of the ship (minimum remodel), bonus applies to all of the ship forms"""

    shipTypes: Optional[list[int]] = Field(alias="shipType", default=None)
    shipNationalities: Optional[list[CountryType]] = Field(
        alias="shipNationality", default=None
    )

    equipmentRequired: Optional[list[EquipmentId]] = Field(
        alias="requires", default=None
    )
    equipmentRequiresLevel: Optional[int] = Field(alias="requiresLevel", default=None)
    """If "EquipmentRequired" is set, equipments requires to be at least this level"""

    numberOfEquipmentsRequired: Optional[int] = Field(alias="requiresNum", default=None)
    """If "EquipmentRequired" is set, you need this number of equipments"""

    equipmentTypesRequired: Optional[list[EquipmentTypes]] = Field(
        alias="requiresType", default=None
    )
    numberOfEquipmentsTypesRequired: Optional[int] = Field(
        alias="requiresNumType", default=None
    )

    equipmentLevel: Optional[int] = Field(alias="level", default=None)
    """Improvement level of the equipment required"""

    numberOfEquipmentsRequiredAfterOtherFilters: Optional[int] = Field(
        alias="num", default=None
    )
    """Number Of Equipments Required after applying the improvement filter"""

    bonuses: FitBonusValue = Field(alias="bonus", default_factory=FitBonusValue)
    """
    Bonuses to apply
    Applied x times, x being the number of equipment matching the conditions of the bonus fit
    If NumberOfEquipmentsRequiredAfterOtherFilters or EquipmentRequired or EquipmentTypesRequired, bonus is applied only once
    """

    bonusesIfSurfaceRadar: Optional[FitBonusValue] = Field(
        alias="bonusSR", default=None
    )
    """Bonuses to apply if ship has a radar with LOS >= 5"""

    bonusesIfAirRadar: Optional[FitBonusValue] = Field(alias="bonusAR", default=None)
    """Bonuses to apply if ship has a radar with AA >= 2"""

    bonusesIfAccuracyRadar: Optional[FitBonusValue] = Field(
        alias="bonusAccR", default=None
    )
    """Bonuses to apply if ship has a radar with ACC >= 8"""


class FitBonusPerEquipment(BaseModel):
    equipmentTypes: Optional[list[EquipmentTypes]] = Field(alias="types", default=None)
    equipmentIds: Optional[list[EquipmentId]] = Field(alias="ids", default=None)
    bonuses: list[FitBonusData] = Field(alias="bonuses")


@dataclass(slots=True)
class FitBonusResult:
    fitBonusData: FitBonusData = std_field(default_factory=FitBonusData)
    equipmentTypes: list[EquipmentTypes] = std_field(default_factory=list)
    equipmentIds: list[EquipmentId] = std_field(default_factory=list)
    fitBonusValue: list[FitBonusValue] = std_field(default_factory=list)

    @property
    def finalBonus(self) -> FitBonusValue:
        result = FitBonusValue.model_construct()

        for v in self.fitBonusValue:
            result += v

        return result
