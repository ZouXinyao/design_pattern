//
// Created by ZXY on 2021/2/16.
//

#include "ConcreteFactory.h"

Car* BenzFactory::CreateCar() {
    return new BenzCar();
}

Ship* BenzFactory::CreateShip() {
    return new BenzShip();
}

Car* BmwFactory::CreateCar() {
    return new BMWCar();
}

Ship* BmwFactory::CreateShip() {
    return new BMWShip();
}