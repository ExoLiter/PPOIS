from PyQt6.QtWidgets import (QDialog, QFormLayout, QLineEdit, 
                             QSpinBox, QPushButton, QVBoxLayout, QTableWidget, 
                             QTableWidgetItem, QLabel, QHBoxLayout, QHeaderView, QAbstractItemView)
from app.models.student import Student

class StudentDialog(QDialog): # Диалог для добавления нового студента
    def __init__(self, parent=None):
        super().__init__(parent) # Инициализация диалога и создание формы для ввода данных студента
        self.setWindowTitle("Добавить студента")
        self.layout = QFormLayout(self) # Создаем форму для ввода данных студента
         
        self.fio_student = QLineEdit()
        self.fio_father = QLineEdit()
        self.income_father = QSpinBox(); self.income_father.setRange(0, 10000000); self.income_father.setSuffix(" ₽")
        self.fio_mother = QLineEdit()
        self.income_mother = QSpinBox(); self.income_mother.setRange(0, 10000000); self.income_mother.setSuffix(" ₽")
        self.brothers = QSpinBox()
        self.sisters = QSpinBox()
        
        self.layout.addRow("ФИО Студента:", self.fio_student)
        self.layout.addRow("ФИО Отца:", self.fio_father)
        self.layout.addRow("Заработок Отца:", self.income_father)
        self.layout.addRow("ФИО Матери:", self.fio_mother)
        self.layout.addRow("Заработок Матери:", self.income_mother)
        self.layout.addRow("Братьев:", self.brothers)
        self.layout.addRow("Сестер:", self.sisters)
        
        btn = QPushButton("Сохранить")
        btn.clicked.connect(self.accept)
        self.layout.addRow(btn)

    def get_data(self):
        return Student(
            self.fio_student.text(), self.fio_father.text(), self.income_father.value(),
            self.fio_mother.text(), self.income_mother.value(),
            self.brothers.value(), self.sisters.value()
        )

class SearchDialog(QDialog):
    def __init__(self, parent=None, manager=None):
        super().__init__(parent)
        self.manager = manager
        self.search_results = []
        self.page_size = 10
        self.current_page = 1

        self.setWindowTitle("Поиск записей (Вариант 10)")
        self.resize(800, 550)
        layout = QVBoxLayout(self)
        
        # --- ФОРМА ПОИСКА ---
        form = QFormLayout()
        
        # 1. ФИО Студента
        self.fio_student = QLineEdit()
        self.fio_student.setPlaceholderText("Имя или фамилия студента")
        
        # 2. ФИО Родителя
        self.parent_fio = QLineEdit()
        self.parent_fio.setPlaceholderText("Имя или фамилия родителя")
        
        # 3. Братья и Сестры
        self.brothers = QSpinBox(); self.brothers.setRange(-1, 20); self.brothers.setValue(-1); self.brothers.setSpecialValueText("Не важно")
        self.sisters = QSpinBox(); self.sisters.setRange(-1, 20); self.sisters.setValue(-1); self.sisters.setSpecialValueText("Не важно")
        
        # 4. Зарплата (Нижняя и верхняя границы)
        self.income_min = QSpinBox(); self.income_min.setRange(0, 10000000)
        self.income_min.setSpecialValueText("Любой")
        self.income_min.setSuffix(" ₽")
        
        self.income_max = QSpinBox(); self.income_max.setRange(0, 10000000); self.income_max.setValue(10000000)
        self.income_max.setSpecialValueText("Без лимита")
        self.income_max.setSuffix(" ₽")
        
        form.addRow("ФИО студента:", self.fio_student)
        form.addRow("ФИО родителя:", self.parent_fio)
        form.addRow("Братьев (точно):", self.brothers)
        form.addRow("Сестер (точно):", self.sisters)
        form.addRow("Доход родителя ОТ:", self.income_min)
        form.addRow("Доход родителя ДО:", self.income_max)
        
        btn = QPushButton("Найти")
        btn.clicked.connect(self.on_search)
        layout.addLayout(form)
        layout.addWidget(btn)
        
        # --- ТАБЛИЦА РЕЗУЛЬТАТОВ ---
        self.table = QTableWidget()
        self.table.setColumnCount(7)
        self.table.setHorizontalHeaderLabels(["Студент", "Отец", "Дох.О", "Мать", "Дох.М", "Бр", "Сес"])
        self.table.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeMode.Stretch)
        self.table.setEditTriggers(QAbstractItemView.EditTrigger.NoEditTriggers)
        layout.addWidget(self.table)

        # --- ПАНЕЛЬ ПАГИНАЦИИ ---
        pag_layout = QHBoxLayout()
        self.btn_prev = QPushButton("< Пред.")
        self.lbl_page = QLabel("Страница: 1/1")
        self.btn_next = QPushButton("След. >")
        
        self.btn_prev.clicked.connect(lambda: self.change_page(-1))
        self.btn_next.clicked.connect(lambda: self.change_page(1))
        
        pag_layout.addStretch()
        pag_layout.addWidget(self.btn_prev)
        pag_layout.addWidget(self.lbl_page)
        pag_layout.addWidget(self.btn_next)
        pag_layout.addStretch()
        layout.addLayout(pag_layout)

    def on_search(self):
        criteria = {}
        if self.fio_student.text(): criteria['fio_student'] = self.fio_student.text()
        if self.parent_fio.text(): criteria['parent_fio'] = self.parent_fio.text()
        if self.brothers.value() >= 0: criteria['brothers'] = self.brothers.value()
        if self.sisters.value() >= 0: criteria['sisters'] = self.sisters.value()
        
        if self.income_min.value() > 0: 
            criteria['income_min'] = self.income_min.value()
        if self.income_max.value() < 10000000: # Если не стоит значение "Без лимита"
            criteria['income_max'] = self.income_max.value()
        
        self.search_results = self.manager.search(criteria)
        self.current_page = 1
        self.update_table()

    def change_page(self, delta):
        total_pages = max(1, (len(self.search_results) + self.page_size - 1) // self.page_size)
        new_page = self.current_page + delta
        if 1 <= new_page <= total_pages:
            self.current_page = new_page
            self.update_table()

    def update_table(self):
        total = len(self.search_results)
        total_pages = max(1, (total + self.page_size - 1) // self.page_size)
        self.lbl_page.setText(f"Страница: {self.current_page}/{total_pages} (Найдено: {total})")

        start = (self.current_page - 1) * self.page_size
        page_data = self.search_results[start : start + self.page_size]

        self.table.setRowCount(0)
        for row, s in enumerate(page_data):
            self.table.insertRow(row)
            self.table.setItem(row, 0, QTableWidgetItem(s.fio_student))
            self.table.setItem(row, 1, QTableWidgetItem(s.fio_father))
            self.table.setItem(row, 2, QTableWidgetItem(str(s.income_father)))
            self.table.setItem(row, 3, QTableWidgetItem(s.fio_mother))
            self.table.setItem(row, 4, QTableWidgetItem(str(s.income_mother)))
            self.table.setItem(row, 5, QTableWidgetItem(str(s.brothers)))
            self.table.setItem(row, 6, QTableWidgetItem(str(s.sisters)))

class DeleteDialog(QDialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setWindowTitle("Удаление записей (Вариант 10)")
        self.resize(400, 300)
        self.layout = QFormLayout(self)
        
        # Поля абсолютно такие же, как в поиске
        self.fio_student = QLineEdit()
        self.fio_student.setPlaceholderText("Например: Иван")
        
        self.parent_fio = QLineEdit()
        self.parent_fio.setPlaceholderText("Например: Петров")
        
        self.brothers = QSpinBox(); self.brothers.setRange(-1, 20); self.brothers.setValue(-1); self.brothers.setSpecialValueText("Не важно")
        self.sisters = QSpinBox(); self.sisters.setRange(-1, 20); self.sisters.setValue(-1); self.sisters.setSpecialValueText("Не важно")
        
        self.income_min = QSpinBox(); self.income_min.setRange(0, 10000000)
        self.income_min.setSpecialValueText("Любой")
        self.income_min.setSuffix(" ₽")
        
        self.income_max = QSpinBox(); self.income_max.setRange(0, 10000000); self.income_max.setValue(10000000)
        self.income_max.setSpecialValueText("Без лимита")
        self.income_max.setSuffix(" ₽")
        
        self.layout.addRow("ФИО студента:", self.fio_student)
        self.layout.addRow("ФИО родителя:", self.parent_fio)
        self.layout.addRow("Братьев (точно):", self.brothers)
        self.layout.addRow("Сестер (точно):", self.sisters)
        self.layout.addRow("Доход от:", self.income_min)
        self.layout.addRow("Доход до:", self.income_max)
        
        btn = QPushButton("Удалить найденных")
        btn.clicked.connect(self.accept)
        self.layout.addRow(btn)

    def get_criteria(self):
        # Собираем критерии так же, как в поиске
        criteria = {}
        if self.fio_student.text(): criteria['fio_student'] = self.fio_student.text()
        if self.parent_fio.text(): criteria['parent_fio'] = self.parent_fio.text()
        if self.brothers.value() >= 0: criteria['brothers'] = self.brothers.value()
        if self.sisters.value() >= 0: criteria['sisters'] = self.sisters.value()
        
        if self.income_min.value() > 0: 
            criteria['income_min'] = self.income_min.value()
        if self.income_max.value() < 10000000:
            criteria['income_max'] = self.income_max.value()
            
        return criteria