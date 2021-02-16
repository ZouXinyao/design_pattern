//
// Created by ZXY on 2021/2/16.
//

#include "ConcreteFactory.h"

Car* BenzFactory::CreateCar() {
    return new Benz();
}

Car* BmwFactory::CreateCar() {
    return new BMW();
}