class Student:
    def __init__(self, fio_student, fio_father, income_father, fio_mother, income_mother, brothers, sisters): 
        self.fio_student = fio_student
        self.fio_father = fio_father
        self.income_father = float(income_father)
        self.fio_mother = fio_mother
        self.income_mother = float(income_mother)
        self.brothers = int(brothers)
        self.sisters = int(sisters)