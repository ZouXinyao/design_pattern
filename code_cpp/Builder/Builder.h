//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_BUILDER_H
#define CODE_CPP_BUILDER_H

#include "Product.h"

class ComputerBuilder {
protected:
    Computer *pComputer{};

public:
    virtual ~ComputerBuilder();

    virtual void BuildCpu() = 0;
    virtual void BuildMainboard() = 0;
    virtual void BuildRam() = 0;
    virtual void BuildVideoCard() = 0;
    virtual Computer* GetComputerResult() = 0;
};

#endif //CODE_CPP_BUILDER_H
