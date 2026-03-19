"""PyQt GUI view."""

from importlib import import_module


def _load_qt_bindings() -> tuple[object, object]:
    candidates = (
        ("PyQt6.QtCore", "PyQt6.QtWidgets"),
        ("PyQt5.QtCore", "PyQt5.QtWidgets"),
    )

    last_error: ImportError | None = None
    for core_module_name, widgets_module_name in candidates:
        try:
            core_module = import_module(core_module_name)
            widgets_module = import_module(widgets_module_name)
            return core_module, widgets_module
        except ImportError as exc:
            last_error = exc

    raise ImportError("PyQt5 or PyQt6 is required for the GUI.") from last_error


QT_CORE, QT_WIDGETS = _load_qt_bindings()
Qt = QT_CORE.Qt
QApplication = QT_WIDGETS.QApplication
QButtonGroup = QT_WIDGETS.QButtonGroup
QCheckBox = QT_WIDGETS.QCheckBox
QComboBox = QT_WIDGETS.QComboBox
QFormLayout = QT_WIDGETS.QFormLayout
QGridLayout = QT_WIDGETS.QGridLayout
QGroupBox = QT_WIDGETS.QGroupBox
QHBoxLayout = QT_WIDGETS.QHBoxLayout
QLabel = QT_WIDGETS.QLabel
QLineEdit = QT_WIDGETS.QLineEdit
QMainWindow = QT_WIDGETS.QMainWindow
QMessageBox = QT_WIDGETS.QMessageBox
QPushButton = QT_WIDGETS.QPushButton
QRadioButton = QT_WIDGETS.QRadioButton
QSpinBox = QT_WIDGETS.QSpinBox
QVBoxLayout = QT_WIDGETS.QVBoxLayout
QWidget = QT_WIDGETS.QWidget


class GuiView(QMainWindow):
    def __init__(self) -> None:
        super().__init__()
        self.setWindowTitle("Модель готовки яичницы")
        self.resize(760, 520)
        self._build_ui()

    def _build_ui(self) -> None:
        central_widget = QWidget()
        main_layout = QVBoxLayout(central_widget)

        form_group = QGroupBox("Параметры")
        form_layout = QFormLayout()

        self.eggs_spin = QSpinBox()
        self.eggs_spin.setRange(1, 10)
        form_layout.addRow("Сколько яиц разбить:", self.eggs_spin)

        stove_row = QHBoxLayout()
        self.power_combo = QComboBox()
        self.power_combo.addItems(["1", "2", "3", "4", "5", "6"])
        self.time_edit = QLineEdit()
        self.time_edit.setPlaceholderText("минуты")
        stove_row.addWidget(QLabel("Мощность"))
        stove_row.addWidget(self.power_combo)
        stove_row.addWidget(QLabel("Время"))
        stove_row.addWidget(self.time_edit)
        form_layout.addRow("Электроплита:", stove_row)

        oil_layout = QHBoxLayout()
        self.oil_group = QButtonGroup(self)
        self.butter_radio = QRadioButton("Сливочное")
        self.sunflower_radio = QRadioButton("Подсолнечное")
        self.butter_radio.setChecked(True)
        self.oil_group.addButton(self.butter_radio)
        self.oil_group.addButton(self.sunflower_radio)
        oil_layout.addWidget(self.butter_radio)
        oil_layout.addWidget(self.sunflower_radio)
        form_layout.addRow("Масло:", oil_layout)

        self.fry_spin = QSpinBox()
        self.fry_spin.setRange(1, 1)
        form_layout.addRow("Сколько яиц обжарить:", self.fry_spin)

        seasonings_layout = QHBoxLayout()
        self.salt_check = QCheckBox("Соль")
        self.pepper_check = QCheckBox("Перец")
        self.paprika_check = QCheckBox("Паприка")
        self.universal_check = QCheckBox("Универсальная")
        seasonings_layout.addWidget(self.salt_check)
        seasonings_layout.addWidget(self.pepper_check)
        seasonings_layout.addWidget(self.paprika_check)
        seasonings_layout.addWidget(self.universal_check)
        form_layout.addRow("Приправы:", seasonings_layout)

        form_group.setLayout(form_layout)

        buttons_group = QGroupBox("Операции")
        buttons_layout = QGridLayout()
        self.break_button = QPushButton("Разбить яйца")
        self.stove_button = QPushButton("Включить плиту")
        self.oil_button = QPushButton("Добавить масло")
        self.fry_button = QPushButton("Обжарить яйца")
        self.season_button = QPushButton("Добавить приправы")
        self.serve_button = QPushButton("Перемешать и подать")
        buttons_layout.addWidget(self.break_button, 0, 0)
        buttons_layout.addWidget(self.stove_button, 0, 1)
        buttons_layout.addWidget(self.oil_button, 0, 2)
        buttons_layout.addWidget(self.fry_button, 1, 0)
        buttons_layout.addWidget(self.season_button, 1, 1)
        buttons_layout.addWidget(self.serve_button, 1, 2)
        buttons_group.setLayout(buttons_layout)

        state_group = QGroupBox("Состояние модели")
        state_layout = QVBoxLayout()
        self.state_label = QLabel()
        self.state_label.setWordWrap(True)
        self.status_label = QLabel("Готово к работе.")
        self.status_label.setWordWrap(True)
        state_layout.addWidget(self.state_label)
        state_layout.addWidget(self.status_label)
        state_group.setLayout(state_layout)

        main_layout.addWidget(form_group)
        main_layout.addWidget(buttons_group)
        main_layout.addWidget(state_group)
        self.setCentralWidget(central_widget)

    def get_break_amount(self) -> int:
        return int(self.eggs_spin.value())

    def get_stove_settings(self) -> tuple[int, int]:
        return int(self.power_combo.currentText()), int(self.time_edit.text().strip())

    def get_oil_type(self) -> str:
        return "сливочное" if self.butter_radio.isChecked() else "подсолнечное"

    def get_fry_amount(self) -> int:
        return int(self.fry_spin.value())

    def get_selected_seasonings(self) -> list[str]:
        items: list[str] = []
        if self.salt_check.isChecked():
            items.append("соль")
        if self.pepper_check.isChecked():
            items.append("перец")
        if self.paprika_check.isChecked():
            items.append("паприка")
        if self.universal_check.isChecked():
            items.append("универсальная")
        return items

    def set_status(self, message: str) -> None:
        self.status_label.setText(message)

    def set_state(self, message: str) -> None:
        self.state_label.setText(message)

    def set_fry_maximum(self, maximum: int) -> None:
        self.fry_spin.setMaximum(max(1, maximum))

    def show_error(self, message: str) -> None:
        QMessageBox.critical(self, "Ошибка", message)
