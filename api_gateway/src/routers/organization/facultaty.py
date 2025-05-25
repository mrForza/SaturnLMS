from fastapi import APIRouter

router = APIRouter(prefix="/facultaties")

@router.post("/")
async def create_facultaty():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_facultaties():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_facultaty():
    return {"message": "Update admin profile"}

@router.delete("/{facultaty_id}")
async def delete_facultaty(facultaty_id: int):
    return {"message": f"Delete facultaty with ID {facultaty_id}"}