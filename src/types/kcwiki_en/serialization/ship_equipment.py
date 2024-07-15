from pydantic import BaseModel, Field


class KcwikiShipEquipment(BaseModel):
    equipment: str | bool = Field(alias="equipment")
    size: int = Field(alias="size")
