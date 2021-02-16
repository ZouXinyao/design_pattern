//
// Created by ZXY on 2021/2/16.
//

#include "ConcreteFactory.h"
#include <iostream>

int main()
{
    Factory *pFactory = Factory::CreateFactory(Factory::FACTORY_TYPE::BENZ_FACTORY);
    Car *pCar = pFactory->CreateCar();
    Ship *pShip = pFactory->CreateShip();

    std::cout << "Benz factory - Car: " << pCar->Name() << std::endl;
    std::cout << "Benz factory - Ship: " << pShip->Name() << std::endl;


    pFactory = Factory::CreateFactory(Factory::FACTORY_TYPE::BMW_FACTORY);
    pCar = pFactory->CreateCar();
    pShip = pFactory->CreateShip();
    std::cout << "Bmw factory - Car: " << pCar->Name() << std::endl;
    std::cout << "Bmw factory - Ship: " << pShip->Name() << std::endl;

    return 0;
}