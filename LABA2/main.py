import sys
from PyQt6.QtWidgets import QApplication
from app.models.manager import StudentManager
from app.views.main_window import MainWindow
from app.controllers.main_controller import MainController

if __name__ == "__main__":
    app = QApplication(sys.argv)
    
    # 1. Берем Модель из файла manager.py
    manager = StudentManager()           
    
    # 2. Берем Контроллер из файла main_controller.py
    controller = MainController(manager) 
    
    # 3. Берем Окно из файла main_window.py
    window = MainWindow(controller)      
    
    # 4. Соединяем
    controller.set_view(window)          
    
    window.show() # Показываем главное окно приложения
    sys.exit(app.exec()) # Запускаем главный цикл приложения (до тех пор, пока пользователь не закроет окно)