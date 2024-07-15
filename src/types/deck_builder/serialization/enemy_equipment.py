from pydantic import BaseModel, Field

from ...equipment_id import EquipmentId


class DeckBuilderEnemyEquipment(BaseModel):
    id: EquipmentId = Field(alias="id")
