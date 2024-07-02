from __future__ import annotations

from pydantic import BaseModel, ConfigDict, Field, field_validator
from typing_extensions import Any, Sequence, Optional

from ..const import AIR_STATE, FleetType, FormationType, SpEffectItemKind
from ..equipment_id import EquipmentId
from ..ship_id import ShipId


class DeckBuilderAirBase(BaseModel):
    name: str = Field(alias="name")
    equipment: DeckBuilderAirBaseEquipmentList = Field(alias="items")
    mode: AIR_STATE = Field(alias="mode")
    distance: int = Field("distance")
    strikePoint: Optional[list[int]] = Field(alias="sp", default=None)


class DeckBuilderAirBaseEquipmentList(BaseModel):
    equipment1: Optional[DeckBuilderEquipment] = Field(alias="i1", default=None)
    equipment2: Optional[DeckBuilderEquipment] = Field(alias="i2", default=None)
    equipment3: Optional[DeckBuilderEquipment] = Field(alias="i3", default=None)
    equipment4: Optional[DeckBuilderEquipment] = Field(alias="i4", default=None)


class DeckBuilderCell(BaseModel):
    cellId: int = Field(alias="c")
    playerFormation: FormationType = Field(alias="pf")
    enemyFormation: FormationType = Field(alias="ef")
    fleet1: DeckBuilderEnemyFleet = Field(alias="f1")
    fleet2: Optional[DeckBuilderEnemyFleet] = Field(alias="f2", default=None)


class DeckBuilderData(BaseModel):
    version: int = Field(alias="version", default=4)
    hqLevel: int = Field(alias="hqlv")
    fleet1: Optional[DeckBuilderFleet] = Field(alias="f1", default=None)
    fleet2: Optional[DeckBuilderFleet] = Field(alias="f2", default=None)
    fleet3: Optional[DeckBuilderFleet] = Field(alias="f3", default=None)
    fleet4: Optional[DeckBuilderFleet] = Field(alias="f4", default=None)
    airBase1: Optional[DeckBuilderAirBase] = Field(alias="a1", default=None)
    airBase2: Optional[DeckBuilderAirBase] = Field(alias="a2", default=None)
    airBase3: Optional[DeckBuilderAirBase] = Field(alias="a3", default=None)
    sortie: Optional[DeckBuilderSortieData] = Field(alias="s", default=None)

    """
    * noro6 will export empty object when no ship in that fleet
    """

    @field_validator("fleet1", "fleet2", "fleet3", "fleet4", mode="before")
    def _emptyDictToNone(v: Any):
        if v == {}:
            return None

        return v


class DeckBuilderEquipment(BaseModel):
    id: EquipmentId = Field(alias="id")
    level: int = Field(alias="rf")
    aircraftLevel: Optional[int] = Field(alias="mas", default=None)


class DeckBuilderEquipmentList(BaseModel):
    equipment1: Optional[DeckBuilderEquipment] = Field(alias="i1", default=None)
    equipment2: Optional[DeckBuilderEquipment] = Field(alias="i2", default=None)
    equipment3: Optional[DeckBuilderEquipment] = Field(alias="i3", default=None)
    equipment4: Optional[DeckBuilderEquipment] = Field(alias="i4", default=None)
    equipment5: Optional[DeckBuilderEquipment] = Field(alias="i5", default=None)
    equipmentExpansion: Optional[DeckBuilderEquipment] = Field(alias="ix", default=None)


class DeckBuilderEnemyEquipment(BaseModel):
    id: EquipmentId = Field(alias="id")


class DeckBuilderEnemyFleet(BaseModel):
    name: Optional[str] = Field(alias="name", default=None)
    ships: list[DeckBuilderEnemyShip] = Field(alias="s")


class DeckBuilderEnemyShip(BaseModel):
    id: ShipId = Field(alias="id")
    equipment: list[DeckBuilderEnemyEquipment] = Field(alias="items")


class DeckBuilderFleet(BaseModel):
    name: str
    type: FleetType = Field(alias="t")
    ship1: Optional[DeckBuilderShip] = Field(alias="s1", default=None)
    ship2: Optional[DeckBuilderShip] = Field(alias="s2", default=None)
    ship3: Optional[DeckBuilderShip] = Field(alias="s3", default=None)
    ship4: Optional[DeckBuilderShip] = Field(alias="s4", default=None)
    ship5: Optional[DeckBuilderShip] = Field(alias="s5", default=None)
    ship6: Optional[DeckBuilderShip] = Field(alias="s6", default=None)
    ship7: Optional[DeckBuilderShip] = Field(alias="s7", default=None)

    model_config = ConfigDict(populate_by_name=True)


class DeckBuilderShip(BaseModel):
    id: ShipId = Field(alias="id")
    level: int = Field(alias="lv")
    isExpansionSlotAvailable: bool = Field(alias="exa")
    equipment: DeckBuilderEquipmentList = Field(alias="items")
    hp: int = Field(alias="hp")
    antiSubmarine: int = Field(alias="asw")
    luck: int = Field(alias="luck")

    # """ Following attribute only provide by 74EOen"""
    firepower: Optional[int] = Field(alias="fp", default=None)
    torpedo: Optional[int] = Field(alias="tp", default=None)
    antiAir: Optional[int] = Field(alias="aa", default=None)
    armor: Optional[int] = Field(alias="ar", default=None)
    evasion: Optional[int] = Field(alias="ev", default=None)
    los: Optional[int] = Field(alias="los", default=None)
    speed: Optional[int] = Field(alias="sp", default=None)
    range: Optional[int] = Field(alias="ra", default=None)
    SpecialEffectItem: Optional[Sequence[DeckBuilderSpecialEffectItem]] = Field(
        alias="spi", default=None
    )

    model_config = ConfigDict(populate_by_name=True)


class DeckBuilderSortieData(BaseModel):
    mapAreaId: int = Field(alias="a")
    mapInfoId: int = Field(alias="i")
    cells: list[DeckBuilderCell] = Field(alias="c")


class DeckBuilderSpecialEffectItem(BaseModel):

    apiKind: SpEffectItemKind = Field(alias="kind")
    """name change from SpEffectItemKind to apiKind, python linter will happy."""

    firepower: int = Field(alias="fp")
    torpedo: int = Field(alias="tp")
    armor: int = Field(alias="ar")
    evasion: int = Field(alias="ev")
