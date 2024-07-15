from pydantic import BaseModel, Field

from .enemy_equipment import DeckBuilderEnemyEquipment
from ...ship_id import ShipId


class DeckBuilderEnemyShip(BaseModel):
    id: ShipId = Field(alias="id")
    equipment: list[DeckBuilderEnemyEquipment] = Field(alias="items")
