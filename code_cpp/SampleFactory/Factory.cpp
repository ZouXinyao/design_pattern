//
// Created by ZXY on 2021/2/16.
//

#include "Factory.h"

Car *Factory::CreateCar(Factory::AUTOMOBLIE_BRAND brand) {
    Car *pCar = nullptr;

    // 简单工厂使用if-else逻辑分支判断
    if(brand == AUTOMOBLIE_BRAND::BENZ_BRAND) pCar = new Benz();
    else if(brand == AUTOMOBLIE_BRAND::BMW_BRAND) pCar = new BMW();

    return pCar;
}
