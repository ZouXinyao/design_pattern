//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_CONCRETEPRODUCTSHIP_H
#define CODE_CPP_CONCRETEPRODUCTSHIP_H

#include "ProductShip.h"

class BenzShip : public Ship
{
public:
    std::string Name() override;
};

class BMWShip : public Ship
{
public:
    std::string Name() override;
};

#endif //CODE_CPP_CONCRETEPRODUCTSHIP_H
