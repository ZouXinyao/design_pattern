//
// Created by ZXY on 2021/2/16.
//

#include "Factory.h"
#include <iostream>

int main()
{
    // 工厂
    auto *pFactory = new Factory();

    // 奔驰汽车
    Car *pCar = pFactory->CreateCar(Factory::BENZ_BRAND);
    std::cout << pCar->Name() << std::endl;

    delete(pCar);

    // 宝马汽车
    pCar = pFactory->CreateCar(Factory::BMW_BRAND);
    std::cout << pCar->Name() << std::endl;

    delete(pCar);

    return 0;
}