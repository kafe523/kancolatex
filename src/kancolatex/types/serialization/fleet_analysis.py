from __future__ import annotations

from pydantic import BaseModel, Field
from typing_extensions import Iterable

from ..equipment_id import EquipmentId
from ..ship_id import ShipId
from ..const import SpEffectItemKind


class FleetAnalysisEquipment(BaseModel):
    id: EquipmentId = Field(alias="api_slotitem_id")
    level: int = Field(alias="api_level")


class FleetAnalysisShip(BaseModel):
    dropId: int = Field(alias="api_id")
    shipId: ShipId = Field(alias="api_ship_id")
    level: int = Field(alias="api_lv")
    modification: Iterable[int] = Field(alias="api_kyouka")
    experience: Iterable[float] = Field(alias="api_exp")
    expansionSlot: int = Field(alias="api_slot_ex")
    sallyArea: int = Field(alias="api_sally_area")
    specialEffectItems: list[FleetAnalysisSpecialEffectItem] = Field(
        alias="api_sp_effect_items"
    )


class FleetAnalysisSpecialEffectItem(BaseModel):
    apiKind: SpEffectItemKind = Field(alias="api_kind")
    firepower: int = Field(alias="api_houg")
    torpedo: int = Field(alias="api_raig")
    armor: int = Field(alias="api_souk")
    evasion: int = Field(alias="api_kaih")
