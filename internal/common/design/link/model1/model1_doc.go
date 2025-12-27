/*
Package model1 提供了责任链设计模式的通用实现

该包实现了责任链模式(model1)，允许将请求沿着处理链传递，直到有处理者处理它。
这种模式将请求的发送者和接收者解耦，使得多个对象都有机会处理请求。

责任链模式 (model1) 包含以下组件:
- ILogicLink: 定义了责任链节点的接口，包含Apply方法用于处理请求
- ILogicChainArmory: 定义了责任链装配的接口，用于构建处理链
- AbstractLogicLink: 提供了责任链节点的抽象实现，包含基本的链式操作
- Model1RuleLogic1/Model1RuleLogic2: 具体的业务逻辑处理器实现
- Model1TradeRuleFactory: 负责装配和创建责任链的工厂类
*/
package model1
