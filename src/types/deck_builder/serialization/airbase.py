from pydantic import BaseModel, Field
from typing_extensions import Optional

from .airbase_equipment import DeckBuilderAirBaseEquipmentList
from ...const import AIR_STATE


class DeckBuilderAirBase(BaseModel):
    name: str = Field(alias="name")
    equipment: DeckBuilderAirBaseEquipmentList = Field(alias="items")
    mode: AIR_STATE = Field(alias="mode")
    distance: int = Field("distance")
    strikePoint: Optional[list[int]] = Field(alias="sp", default=None)
