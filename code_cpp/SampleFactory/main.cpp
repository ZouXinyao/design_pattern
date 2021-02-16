//
// Created by ZXY on 2021/2/16.
//

#include "Factory.h"
#include <iostream>

int main()
{
    // ����
    auto *pFactory = new Factory();

    // ��������
    Car *pCar = pFactory->CreateCar(Factory::BENZ_BRAND);
    std::cout << pCar->Name() << std::endl;

    delete(pCar);

    // ��������
    pCar = pFactory->CreateCar(Factory::BMW_BRAND);
    std::cout << pCar->Name() << std::endl;

    delete(pCar);

    return 0;
}