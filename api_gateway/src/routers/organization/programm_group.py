from fastapi import APIRouter

router = APIRouter(prefix="/programm_groups")

@router.post("/")
async def create_programm_group():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_programm_groups():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_programm_group():
    return {"message": "Update admin profile"}

@router.delete("/{programm_group_id}")
async def delete_programm_group(programm_group_id: int):
    return {"message": f"Delete programm_group with ID {programm_group_id}"}