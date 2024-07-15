from pydantic import BaseModel, Field
from typing_extensions import Optional

from ...equipment_id import EquipmentId


class DeckBuilderEquipment(BaseModel):
    id: EquipmentId = Field(alias="id")
    level: int = Field(alias="rf")
    aircraftLevel: Optional[int] = Field(alias="mas", default=None)
