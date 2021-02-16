//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_DIRECTOR_H
#define CODE_CPP_DIRECTOR_H

#include "Builder.h"

class Director {
public:
    void CreateComputer(ComputerBuilder* builder);
};


#endif //CODE_CPP_DIRECTOR_H
