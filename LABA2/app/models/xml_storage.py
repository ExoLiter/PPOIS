import xml.dom.minidom
import xml.sax
from app.models.student import Student

class StudentHandler(xml.sax.ContentHandler):
    """SAX Парсер для чтения"""
    def __init__(self):
        self.students = [] # Список для хранения загруженных студентов
        self.current_data = "" # Временная переменная для хранения текущих данных элемента
        self.student_data = {} # Словарь для хранения данных текущего студента во время парсинга

    def startElement(self, tag, attrs):
        self.current_data = ""
        if tag == "student":
            self.student_data = {}

    def characters(self, content):
        self.current_data += content

    def endElement(self, tag):
        if tag == "student":
            s = Student(
                self.student_data.get('fio_student', ''),
                self.student_data.get('fio_father', ''),
                self.student_data.get('income_father', 0),
                self.student_data.get('fio_mother', ''),
                self.student_data.get('income_mother', 0),
                self.student_data.get('brothers', 0),
                self.student_data.get('sisters', 0)
            )
            self.students.append(s)
        elif tag in ['fio_student', 'fio_father', 'income_father', 'fio_mother', 'income_mother', 'brothers', 'sisters']:
            self.student_data[tag] = self.current_data

class XmlStorage:
    @staticmethod
    def save(filename, students):
        """DOM Парсер для записи"""
        doc = xml.dom.minidom.Document()
        root = doc.createElement('students_database')
        doc.appendChild(root)

        for s in students:
            item = doc.createElement('student')
            
            def create_node(name, value):
                node = doc.createElement(name)
                text = doc.createTextNode(str(value))
                node.appendChild(text)
                return node

            item.appendChild(create_node('fio_student', s.fio_student))
            item.appendChild(create_node('fio_father', s.fio_father))
            item.appendChild(create_node('income_father', s.income_father))
            item.appendChild(create_node('fio_mother', s.fio_mother))
            item.appendChild(create_node('income_mother', s.income_mother))
            item.appendChild(create_node('brothers', s.brothers))
            item.appendChild(create_node('sisters', s.sisters))
            
            root.appendChild(item)

        with open(filename, "w", encoding="utf-8") as f:
            f.write(doc.toprettyxml(indent="  "))

    @staticmethod
    def load(filename):
        handler = StudentHandler()
        parser = xml.sax.make_parser()
        parser.setContentHandler(handler)
        parser.parse(filename)
        return handler.students