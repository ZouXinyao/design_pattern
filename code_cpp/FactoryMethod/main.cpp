//
// Created by ZXY on 2021/2/16.
//

#include "ConcreteFactory.h"
#include <iostream>

int main()
{
    Factory *pFactory = new BenzFactory();
    Car *pCar = pFactory->CreateCar();
    std::cout << "Benz factory: " << pCar->Name() << std::endl;


    pFactory = new BmwFactory();
    pCar = pFactory->CreateCar();
    std::cout << "Bmw factory: " << pCar->Name() << std::endl;

    return 0;
}