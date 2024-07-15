from __future__ import annotations

from dataclasses import dataclass, field

from . import FitBonusData, FitBonusValue
from .. import EquipmentId
from ..const import EquipmentTypes

@dataclass(slots=True)
class FitBonusResult:
    fitBonusData: FitBonusData = field(default_factory=FitBonusData)
    equipmentTypes: list[EquipmentTypes] = field(default_factory=list)
    equipmentIds: list[EquipmentId] = field(default_factory=list)
    fitBonusValue: list[FitBonusValue] = field(default_factory=list)

    @property
    def finalBonus(self) -> FitBonusValue:
        result = FitBonusValue.model_construct()

        for v in self.fitBonusValue:
            result += v

        return result
