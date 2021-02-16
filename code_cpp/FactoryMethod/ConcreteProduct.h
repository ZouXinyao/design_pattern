//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_CONCRETEPRODUCT_H
#define CODE_CPP_CONCRETEPRODUCT_H

#include "Product.h"

class Benz : public Car
{
public:
    std::string Name() override;
};

class BMW : public Car
{
public:
    std::string Name() override;
};


#endif //CODE_CPP_CONCRETEPRODUCT_H
