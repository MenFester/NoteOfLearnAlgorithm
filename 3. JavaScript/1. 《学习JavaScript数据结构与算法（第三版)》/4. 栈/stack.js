class Stack {
    constructor() {
        this.count = 0;
        this.items = {};
    }

    push(element) {
        this.items[this.count] = element;
        this.count++;
    }

    pop() {
        if (this.isEmpty()) {
            return undefined;
        }

        this.count--;
        const result = this.items[this.count];
        delete this.items[this.count];    // 使用JavaScript对象的delete运算符删除对象的一个特定的值
        return result;
    }

    size() {
        return this.count;
    }

    isEmpty() {
        return this.count == 0;
    }

    peek() {
        if (this.isEmpty()) {
            return undefined;
        }

        return this.items[this.count - 1];
    }

    clear() {
        this.items = {};
        this.count = 0;
    }

    toString() {
        if (this.isEmpty()) {
            return '"stack is empty"';
        }

        let objString = `${this.items[0]}`;
        for (let i = 1; i < this.count; i++) {
            objString = `${objString}, ${this.items[i]}`;
        }
        return objString;
    }
}

module.exports = {Stack};

const stack = new Stack();
stack.push(5);
stack.push(8);
console.log(stack.isEmpty());
console.log(stack.size());

stack.push(11);
stack.push(12);
console.log(stack.toString());
console.log(stack.pop());
console.log(stack.toString());

console.log(stack.peek());

stack.clear();
console.log(stack.isEmpty());
console.log(stack.size());
console.log(stack.toString());

console.log(Object.getOwnPropertyNames(stack));