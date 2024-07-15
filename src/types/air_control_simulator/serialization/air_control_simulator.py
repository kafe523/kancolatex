from pydantic import BaseModel, Field
from typing_extensions import Iterable, Optional

from ...deck_builder import DeckBuilderData
from ...fleet_analysis import FleetAnalysisShip, FleetAnalysisEquipment


class AirControlSimulator(BaseModel):
    fleet: Optional[DeckBuilderData] = Field(alias="predeck")
    ships: Optional[Iterable[FleetAnalysisShip]] = Field(alias="ships")
    equipment: Optional[Iterable[FleetAnalysisEquipment]] = Field(alias="items")
