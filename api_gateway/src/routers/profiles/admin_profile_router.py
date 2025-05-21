from fastapi import APIRouter

router = APIRouter(prefix="/admin-profiles")

@router.post("/")
async def create_admin_profile():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_admin_profiles():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_admin_profile():
    return {"message": "Update admin profile"}

@router.delete("/{profile_id}")
async def delete_admin_profile(profile_id: int):
    return {"message": f"Delete admin profile with ID {profile_id}"}