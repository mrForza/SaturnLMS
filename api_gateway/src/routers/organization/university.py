from fastapi import APIRouter

router = APIRouter(prefix="/universities")

@router.post("/")
async def create_university():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_universities():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_university():
    return {"message": "Update admin profile"}

@router.delete("/{university_id}")
async def delete_university(university_id: int):
    return {"message": f"Delete university with ID {university_id}"}