from fastapi import APIRouter

router = APIRouter(prefix="/programms")

@router.post("/")
async def create_programm():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_programms():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_programm():
    return {"message": "Update admin profile"}

@router.delete("/{programm_id}")
async def delete_programm(programm_id: int):
    return {"message": f"Delete programm with ID {programm_id}"}