from pydantic import BaseModel, Field
from typing_extensions import Optional

from . import FitBonusData
from ... import EquipmentId
from ...const import EquipmentTypes

class FitBonusPerEquipment(BaseModel):
    equipmentTypes: Optional[list[EquipmentTypes]] = Field(alias="types", default=None)
    equipmentIds: Optional[list[EquipmentId]] = Field(alias="ids", default=None)
    bonuses: list[FitBonusData] = Field(alias="bonuses")
