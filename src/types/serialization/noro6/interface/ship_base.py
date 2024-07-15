from abc import ABC
from dataclasses import dataclass, InitVar

from typing_extensions import Any

from ..aerial_combat.anti_air_cut_in import AntiAirCutIn
from ..enemy.enemy_master import EnemyMaster
from ..fleet.ship_master import ShipMaster
from ..item.item import Item
from ..item.item_bonus import ItemBonusStatus


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
