from typing_extensions import MutableSequence

from ...eo_en.fit_bonus import FitBonusValue as ItemBonusStatus
from .....database import DATABASE


def getTotalBonus(bonuses: MutableSequence[ItemBonusStatus]) -> ItemBonusStatus:
    """装備ボーナス配列から合計のボーナスを算出する"""

    bonus = ItemBonusStatus()

    for v in bonuses:
        bonus += v

    return bonus


def bonusData():
    for v in DATABASE.QueryFitBonusAllFromEOEn():
        yield v
