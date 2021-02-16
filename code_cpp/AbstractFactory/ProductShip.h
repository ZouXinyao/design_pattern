//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_PRODUCTSHIP_H
#define CODE_CPP_PRODUCTSHIP_H

#include <string>

class Ship
{
public:
    virtual std::string Name() = 0;

    virtual ~Ship() = 0;
};

#endif //CODE_CPP_PRODUCTSHIP_H
