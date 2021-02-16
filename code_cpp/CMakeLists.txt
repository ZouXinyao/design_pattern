cmake_minimum_required(VERSION 3.17)
project(code_cpp)

set(CMAKE_C_STANDARD 11)

include_directories(.)
include_directories(SampleFactory)

# 简单工厂模式
#add_executable(code_cpp
#        SampleFactory/ConcreteProduct.cpp
#        SampleFactory/Product.cpp
#        SampleFactory/Factory.cpp
#        SampleFactory/main.cpp)

# 工厂方法模式
#add_executable(code_cpp
#        FactoryMethod/ConcreteProduct.cpp
#        FactoryMethod/ConcreteFactory.cpp
#        FactoryMethod/Product.cpp
#        FactoryMethod/Factory.cpp
#        FactoryMethod/main.cpp)

# 抽象工厂模式
#add_executable(code_cpp
#        AbstractFactory/ConcreteProductCar.cpp
#        AbstractFactory/ConcreteProductShip.cpp
#        AbstractFactory/ConcreteFactory.cpp
#        AbstractFactory/ProductCar.cpp
#        AbstractFactory/ProductShip.cpp
#        AbstractFactory/Factory.cpp
#        AbstractFactory/main.cpp)

# 单例模式
#add_executable(code_cpp
#        Singleton/Singleton.cpp
#        Singleton/main.cpp)

# 建造者模式
add_executable(code_cpp
        Builder/Builder.cpp
        Builder/ConcreteBuilder.cpp
        Builder/Director.cpp
        Builder/Product.cpp
        Builder/main.cpp)