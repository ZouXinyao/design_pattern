//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_CONCRETEBUILDER_H
#define CODE_CPP_CONCRETEBUILDER_H

#include "Builder.h"

class LenovoBuilder : public ComputerBuilder
{
public:
    LenovoBuilder()       { pComputer = new Computer(); }
    void BuildCpu()       override { pComputer->SetCpu("i7-10700k"); }
    void BuildMainboard() override { pComputer->SetMainboard("Intel DH57DD"); }
    void BuildRam()       override { pComputer->SetRam("DDR4"); }
    void BuildVideoCard() override { pComputer->SetVideoCard("NVIDIA Geforce 920MX"); }
    Computer* GetComputerResult() override { return pComputer; }

};

class DellBuilder : public ComputerBuilder
{
public:
    DellBuilder()         { pComputer = new Computer(); }
    void BuildCpu()       override { pComputer->SetCpu("i7-7500U"); }
    void BuildMainboard() override { pComputer->SetMainboard("Intel DP55KG"); }
    void BuildRam()       override { pComputer->SetRam("DDR5"); }
    void BuildVideoCard() override { pComputer->SetVideoCard("NVIDIA GeForce 940MX"); }
    Computer* GetComputerResult() override { return pComputer; }
};


#endif //CODE_CPP_CONCRETEBUILDER_H
