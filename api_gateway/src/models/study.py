from pydantic import BaseModel
from uuid import UUID
from typing import List


class Course(BaseModel):
    id: UUID
    name: str
    description: str
    formula: str
    languages: List[str]
    teachers: List[UUID]
    students: List[UUID]
    lessons: List[UUID]


class Homework(BaseModel):
    id: UUID
    name: str
    description: str
    files: List[UUID]


class File(BaseModel):
    bucket_id: UUID
    name: str
    extension: str


class Lesson(BaseModel):
    id: UUID
    name: str
    description: str
    type: bool
    files: List[File]
    homework: UUID