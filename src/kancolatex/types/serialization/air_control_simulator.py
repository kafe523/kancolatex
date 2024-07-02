from pydantic import BaseModel, Field
from typing_extensions import Iterable, Optional

from .fleet_analysis import FleetAnalysisEquipment, FleetAnalysisShip
from .deck_builder import DeckBuilderData


class AirControlSimulator(BaseModel):
    fleet: Optional[DeckBuilderData] = Field(alias="predeck")
    ships: Optional[Iterable[FleetAnalysisShip]] = Field(alias="ships")
    equipment: Optional[Iterable[FleetAnalysisEquipment]] = Field(alias="items")
