@echo off

:: Create directories
mkdir "api/user/service"
mkdir "api/user/controller"
mkdir "api/user/repository"
mkdir "api/user/entity"
mkdir "api/user/dto"

mkdir "api/note/service"
mkdir "api/note/controller"
mkdir "api/note/repository"
mkdir "api/note/entity"
mkdir "api/note/dto"

mkdir "api/food/service"
mkdir "api/food/controller"
mkdir "api/food/repository"
mkdir "api/food/entity"
mkdir "api/food/dto"

mkdir "api/invoice/service"
mkdir "api/invoice/controller"
mkdir "api/invoice/repository"
mkdir "api/invoice/entity"
mkdir "api/invoice/dto"

mkdir "api/menu/service"
mkdir "api/menu/controller"
mkdir "api/menu/repository"
mkdir "api/menu/entity"
mkdir "api/menu/dto"

mkdir "api/order/service"
mkdir "api/order/controller"
mkdir "api/order/repository"
mkdir "api/order/entity"
mkdir "api/order/dto"

mkdir "api/table/service"
mkdir "api/table/controller"
mkdir "api/table/repository"
mkdir "api/table/entity"
mkdir "api/table/dto"

:: Create files in directories
echo. > "api/user/service/user.service.go"
echo. > "api/user/controller/user.controller.go"
echo. > "api/user/repository/user.repository.go"
echo. > "api/user/entity/user.entity.go"
echo. > "api/user/dto/user.dto.go"
echo. > "api/user/routes.go"

echo. > "api/note/service/note.service.go"
echo. > "api/note/controller/note.controller.go"
echo. > "api/note/repository/note.repository.go"
echo. > "api/note/entity/note.entity.go"
echo. > "api/note/dto/note.dto.go"
echo. > "api/note/note.routes.go"

echo. > "api/food/service/food.service.go"
echo. > "api/food/controller/food.controller.go"
echo. > "api/food/repository/food.repository.go"
echo. > "api/food/entity/food.entity.go"
echo. > "api/food/dto/food.dto.go"
echo. > "api/food/food.routes.go"

echo. > "api/invoice/service/invoice.service.go"
echo. > "api/invoice/controller/invoice.controller.go"
echo. > "api/invoice/repository/invoice.repository.go"
echo. > "api/invoice/entity/invoice.entity.go"
echo. > "api/invoice/dto/invoice.dto.go"
echo. > "api/invoice/invoice.routes.go"

echo. > "api/menu/service/menu.service.go"
echo. > "api/menu/controller/menu.controller.go"
echo. > "api/menu/repository/menu.repository.go"
echo. > "api/menu/entity/menu.entity.go"
echo. > "api/menu/dto/menu.dto.go"
echo. > "api/menu/menu.routes.go"

echo. > "api/order/service/order.service.go"
echo. > "api/order/controller/order.controller.go"
echo. > "api/order/repository/order.repository.go"
echo. > "api/order/entity/order.entity.go"
echo. > "api/order/entity/orderItem.entity.go"
echo. > "api/order/dto/order.dto.go"
echo. > "api/order/order.routes.go"

echo. > "api/table/service/table.service.go"
echo. > "api/table/controller/table.controller.go"
echo. > "api/table/repository/table.repository.go"
echo. > "api/table/entity/table.entity.go"
echo. > "api/table/dto/table.dto.go"
echo. > "api/table/table.routes.go"

echo. > "api/main.go"
echo. > "api/routes.go"
echo. > "api/loadEnv.go"
echo. > "Makefile"


echo "Package main has been created successfully."
