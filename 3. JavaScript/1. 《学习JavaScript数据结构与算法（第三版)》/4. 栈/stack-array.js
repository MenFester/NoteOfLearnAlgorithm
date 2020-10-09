class Stack {
    constructor() {
        this.items = [];
    }

    push(element) {
        this.items.push(element);    // 用数组的push方法实现
    }

    pop() {
        return this.items.pop();
    }

    peek() {
        return this.items[this.items.length - 1];
    }

    isEmpty() {
        return this.items.length == 0;    // 数组没有isEmpty方法
    }

    size() {
        return this.items.length;
    }

    clear() {
        this.items = [];
    }
}

const stack = new Stack();
console.log(stack.isEmpty());

stack.push(5);
stack.push("a");

console.log(stack.peek());

stack.push(11);
console.log(stack.size());
console.log(stack.isEmpty());

stack.push(15);
stack.pop();
stack.pop();
console.log(stack.size());