//
// Created by ZXY on 2021/2/16.
//

#include "Factory.h"
#include "ConcreteFactory.h"

Factory* Factory::CreateFactory(FACTORY_TYPE factoryType)
{
    Factory *pFactory = nullptr;
    if(factoryType == FACTORY_TYPE::BENZ_FACTORY) pFactory = new BenzFactory();
    else if(factoryType == FACTORY_TYPE::BMW_FACTORY) pFactory = new BmwFactory();

    return pFactory;
}