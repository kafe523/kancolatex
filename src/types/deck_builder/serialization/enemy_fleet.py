from pydantic import BaseModel, Field
from typing_extensions import Optional

from .enemy_ship import DeckBuilderEnemyShip


class DeckBuilderEnemyFleet(BaseModel):
    name: Optional[str] = Field(alias="name", default=None)
    ships: list[DeckBuilderEnemyShip] = Field(alias="s")
