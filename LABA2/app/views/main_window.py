from PyQt6.QtWidgets import (QMainWindow, QWidget, QVBoxLayout, QTableWidget, 
                             QTableWidgetItem, QHeaderView, QHBoxLayout, QPushButton, 
                             QLabel)
from PyQt6.QtGui import QAction

class MainWindow(QMainWindow):
    def __init__(self, controller):
        super().__init__()
        self.controller = controller # Ссылка на контроллер для связи
        self.setWindowTitle("Лаб. работа №2 - MVC")
        self.resize(1000, 600)
        
        central = QWidget()
        self.setCentralWidget(central)
        self.layout = QVBoxLayout(central)
        
        self._create_ui()

    def _create_ui(self):
        # Меню и Тулбар
        menubar = self.menuBar()
        file_menu = menubar.addMenu("Файл")
        
        # Actions
        self.act_add = QAction("Добавить", self)
        self.act_add.triggered.connect(self.controller.add_record)
        
        self.act_save = QAction("Сохранить", self)
        self.act_save.triggered.connect(self.controller.save_file)
        
        self.act_load = QAction("Загрузить", self)
        self.act_load.triggered.connect(self.controller.load_file)

        self.act_search = QAction("Поиск", self)
        self.act_search.triggered.connect(self.controller.open_search)
        
        self.act_del = QAction("Удалить", self)
        self.act_del.triggered.connect(self.controller.open_delete)

        self.act_gen = QAction("Генерация (50)", self)
        self.act_gen.triggered.connect(self.controller.gen_data)

        file_menu.addAction(self.act_save)
        file_menu.addAction(self.act_load)
        file_menu.addAction(self.act_gen)
        
        toolbar = self.addToolBar("Main")
        toolbar.addAction(self.act_add)
        toolbar.addAction(self.act_search)
        toolbar.addAction(self.act_del)

        # Таблица
        self.table = QTableWidget()
        self.table.setColumnCount(7)
        self.table.setHorizontalHeaderLabels(["Студент", "Отец", "Дох.О", "Мать", "Дох.М", "Бр", "Сес"])
        self.table.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeMode.Stretch)
        self.layout.addWidget(self.table)
        
        # Пагинация (кнопки и лейбл)
        pag_layout = QHBoxLayout()
        self.btn_prev = QPushButton("<")
        self.btn_next = QPushButton(">")
        self.lbl_page = QLabel("1/1")
        self.btn_prev.clicked.connect(lambda: self.controller.change_page('prev'))
        self.btn_next.clicked.connect(lambda: self.controller.change_page('next'))
        
        pag_layout.addWidget(self.btn_prev)
        pag_layout.addWidget(self.lbl_page)
        pag_layout.addWidget(self.btn_next)
        self.layout.addLayout(pag_layout)

    def update_table(self, data):
        self.table.setRowCount(0)
        for row, s in enumerate(data):
            self.table.insertRow(row)
            self.table.setItem(row, 0, QTableWidgetItem(s.fio_student))
            self.table.setItem(row, 1, QTableWidgetItem(s.fio_father))
            self.table.setItem(row, 2, QTableWidgetItem(str(s.income_father)))
            self.table.setItem(row, 3, QTableWidgetItem(s.fio_mother))
            self.table.setItem(row, 4, QTableWidgetItem(str(s.income_mother)))
            self.table.setItem(row, 5, QTableWidgetItem(str(s.brothers)))
            self.table.setItem(row, 6, QTableWidgetItem(str(s.sisters)))

    def update_pagination_label(self, current, total):
        self.lbl_page.setText(f"{current}/{total}")