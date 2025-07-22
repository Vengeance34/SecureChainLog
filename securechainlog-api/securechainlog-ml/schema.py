from pydantic import BaseModel

class AssetLogInput(BaseModel):
    asset_id: int
    status: str
    timestamp: str  # or datetime if needed
