import random
from app.models.student import Student
from app.models.xml_storage import XmlStorage

class StudentManager:
    def __init__(self):
        self.students = []
    
    def add_student(self, student):
        self.students.append(student)

    def get_all(self):
        return self.students

    def save_to_file(self, filename):
        XmlStorage.save(filename, self.students)

    def load_from_file(self, filename):
        self.students = XmlStorage.load(filename)

    def _check_conditions(self, student, criteria):
        # 1. Поиск по ФИО студента
        if criteria.get('fio_student') and criteria['fio_student'].lower() not in student.fio_student.lower():
            return False
        
        # 2. Поиск по ФИО родителя
        if criteria.get('parent_fio'):
            p_fio = criteria['parent_fio'].lower()
            if (p_fio not in student.fio_father.lower()) and (p_fio not in student.fio_mother.lower()):
                return False
        
        # 3. Братья и сестры (учитываем, что 0 это тоже допустимое значение)
        if criteria.get('brothers') is not None and criteria['brothers'] >= 0:
             if student.brothers != criteria['brothers']:
                 return False
        if criteria.get('sisters') is not None and criteria['sisters'] >= 0:
             if student.sisters != criteria['sisters']:
                 return False

        # 4. Доходы
        min_inc = criteria.get('income_min', 0)
        max_inc = criteria.get('income_max', float('inf'))
        
        father_ok = (min_inc <= student.income_father <= max_inc)
        mother_ok = (min_inc <= student.income_mother <= max_inc)
        
        if not (father_ok or mother_ok):
            return False

        return True

    def search(self, criteria):
        return [s for s in self.students if self._check_conditions(s, criteria)]

    def delete(self, criteria):
        # Универсальное удаление (работает по любым критериям)
        initial_len = len(self.students)
        self.students = [s for s in self.students if not self._check_conditions(s, criteria)]
        return initial_len - len(self.students)

    def generate_fake_data(self, count=50):
        # Умная генерация (учитывает пол и лимит повторений фамилий)
        male_names = ["Иван", "Алексей", "Дмитрий", "Максим", "Сергей", "Михаил", "Андрей"]
        female_names = ["Анна", "Мария", "Ольга", "Елена", "Екатерина", "Дарья", "Наталья"]
        
        base_surnames = ["Иванов", "Петров", "Сидоров", "Смирнов", "Кузнецов", "Попов", "Соколов", "Лебедев", 
                         "Козлов", "Новиков", "Морозов", "Волков", "Алексеев", "Зайцев", "Егоров", "Николаев", 
                         "Власов", "Степанов", "Макаров", "Орлов"] 

        surname_counts = {surname: 0 for surname in base_surnames}

        for _ in range(count):
            # Выбираем только те фамилии, которые использовались меньше 3 раз
            available_surnames = [s for s, c in surname_counts.items() if c < 3]
            if not available_surnames: 
                break 
            
            base_surname = random.choice(available_surnames)
            surname_counts[base_surname] += 1

            is_male_student = random.choice([True, False])
            
            # Добавляем "а" для девочек
            if is_male_student:
                stud_fio = f"{base_surname} {random.choice(male_names)}"
            else:
                stud_fio = f"{base_surname}а {random.choice(female_names)}"

            father_fio = f"{base_surname} {random.choice(male_names)}"
            mother_fio = f"{base_surname}а {random.choice(female_names)}"

            s = Student(
                stud_fio, 
                father_fio, 
                random.randint(30, 150) * 1000,
                mother_fio, 
                random.randint(30, 150) * 1000,
                random.randint(0, 3), 
                random.randint(0, 3)
            )
            self.students.append(s)