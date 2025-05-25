from pydantic import BaseModel
from uuid import UUID
from datetime import datetime
from typing import Optional


class SystemUserProfile(BaseModel):
    id: UUID
    email: str
    password: str
    last_login: datetime
    last_active: datetime
    registration_time: datetime


class UserProfile(BaseModel):
    id: UUID
    first_name: str
    last_name: str
    father_name: str
    age: int
    gender: str
    about_me: str
    interests: str
    student_profile: Optional['StudentProfile'] = None
    teacher_profile: Optional['TeacherProfile'] = None
    administrator_profile: Optional['AdminProfile'] = None


class StudentProfile(BaseModel):
    id: UUID
    university_name: str
    facultaty_name: str
    program_name: str
    group_number: int
    course_number: int
    user_profile_id: UUID


class TeacherProfile(BaseModel):
    id: UUID
    education: str
    scientific_experience: str
    teaching_experience: str
    professional_interests: str
    achievements: str
    languages: str
    user_profile_id: UUID


class AdminProfile(BaseModel):
    id: UUID
    education: str
    work_experience: str
    achievements: str
    languages: str
    user_profile_id: UUID
