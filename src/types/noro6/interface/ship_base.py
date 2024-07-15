from __future__ import annotations

from abc import ABC
from dataclasses import dataclass, InitVar

from typing_extensions import Any, TYPE_CHECKING

if TYPE_CHECKING:
    from ..aerial_combat import AntiAirCutIn
    from ..item import Item, ItemBonusStatus

    from ..enemy import EnemyMaster
    from ..fleet import ShipMaster


@dataclass(slots=True)
class ShipBase(ABC):
    builder: InitVar[Any]

    data: ShipMaster | EnemyMaster
    items: list[Item]
    exItem: Item
    isEscort: bool
    antiAir: int
    antiAirBonus: float
    antiAirCutIn: list[AntiAirCutIn]
    specialKokakuCount: int
    kokakuCount: int
    specialKijuCount: int
    kijuCount: int
    antiAirRadarCount: int
    surfaceRadarCount: int
    koshaCount: int
    hp: int
    itemBonusStatus: ItemBonusStatus
    enabledAircraftNightAttack: bool
