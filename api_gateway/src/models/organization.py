from pydantic import BaseModel
from uuid import UUID


class University(BaseModel):
    name: str
    description: str
    legal_address: str
    actual_address: str
    inn: str
    owner_id: UUID


class Facultaty(BaseModel):
    name: str
    description: str
    university_name: str


class Program(BaseModel):
    name: str
    description: str
    type: str
    languages: str
    facultaty_name: str


class ProgramGroup(BaseModel):
    number: int
    name: str
    course_number: int
    program_name: str
