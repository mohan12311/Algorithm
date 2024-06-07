class MinHeap {
    constructor() {
        this.heap = [];
    }

    push(val) {
        this.heap.push(val);
        this.bubbleUp();
    }

    pop() {
        if (this.heap.length === 1) {
            return this.heap.pop();
        }
        const top = this.heap[0];
        this.heap[0] = this.heap.pop();
        this.bubbleDown();
        return top;
    }

    bubbleUp() {
        let index = this.heap.length - 1;
        while (index > 0) {
            const element = this.heap[index];
            const parentIndex = Math.floor((index - 1) / 2);
            const parent = this.heap[parentIndex];

            if (parent.distance <= element.distance) break;

            this.heap[index] = parent;
            this.heap[parentIndex] = element;
            index = parentIndex;
        }
    }

    bubbleDown() {
        let index = 0;
        const length = this.heap.length;
        const element = this.heap[0];

        while (true) {
            const leftChildIndex = 2 * index + 1;
            const rightChildIndex = 2 * index + 2;
            let leftChild, rightChild;
            let swap = null;

            if (leftChildIndex < length) {
                leftChild = this.heap[leftChildIndex];
                if (leftChild.distance < element.distance) {
                    swap = leftChildIndex;
                }
            }

            if (rightChildIndex < length) {
                rightChild = this.heap[rightChildIndex];
                if (
                    (swap === null && rightChild.distance < element.distance) ||
                    (swap !== null && rightChild.distance < leftChild.distance)
                ) {
                    swap = rightChildIndex;
                }
            }

            if (swap === null) break;

            this.heap[index] = this.heap[swap];
            this.heap[swap] = element;
            index = swap;
        }
    }

    isEmpty() {
        return this.heap.length === 0;
    }
}

class Graph {
    constructor() {
        this.nodes = new Map();
    }

    addNode(n) {
        if (!this.nodes.has(n)) {
            this.nodes.set(n, []);
        }
    }

    addEdge(start, end, cost) {
        this.addNode(start);
        this.addNode(end);
        this.nodes.get(start).push({ linked: end, cost });
        this.nodes.get(end).push({ linked: start, cost });
    }

    dijkstra(start) {
        const distances = new Map();
        const visited = new Set();
        const heap = new MinHeap();

        this.nodes.forEach((_, key) => {
            distances.set(key, Infinity);
        });
        distances.set(start, 0);
        heap.push({ node: start, distance: 0 });

        while (!heap.isEmpty()) {
            const { node: current, distance: currentDistance } = heap.pop();

            if (visited.has(current)) continue;
            visited.add(current);

            for (let i = 0; i < this.nodes.get(current).length; i++) {
                const now = this.nodes.get(current)[i];
                const { linked, cost } = now;
                const distance = currentDistance + cost;

                if (distance < distances.get(linked)) {
                    distances.set(linked, distance);
                    heap.push({ node: linked, distance });
                }
            }
        }

        return distances;
    }
}

function solution(n, s, a, b, fares) {
    const graph = new Graph();

    fares.forEach(([start, end, cost]) => {
        graph.addEdge(start, end, cost);
    });

    const startDistances = graph.dijkstra(s);
    const aDistances = graph.dijkstra(a);
    const bDistances = graph.dijkstra(b);

    let minFare = Infinity;
    for (let i = 1; i <= n; i++) {
        const totalFare = startDistances.get(i) + aDistances.get(i) + bDistances.get(i);
        if (totalFare < minFare) {
            minFare = totalFare;
        }
    }

    return minFare;
}