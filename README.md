# Golang Algorithms and Data Structures and Go Solution for LeetCode algorithm problems

[![Build Status](https://travis-ci.org/TomorrowWu/golang-algorithms.svg?branch=master)](https://travis-ci.org/TomorrowWu/golang-algorithms)
[![codecov](https://codecov.io/gh/TomorrowWu/golang-algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/TomorrowWu/golang-algorithms)

This repository contains Golang based examples of many
popular algorithms and data structures.

Each algorithm and data structure has its own separate README
with related explanations and links for further reading (including ones
to videos).

_Read this in other languages:_
[_简体中文_](README.zh-CN.md)

## Data Structures

A data structure is a particular way of organizing and storing data in a computer so that it can
be accessed and modified efficiently. More precisely, a data structure is a collection of data
values, the relationships among them, and the functions or operations that can be applied to
the data.

`B` - Beginner, `A` - Advanced

* `B` [Linked List](./data-structures/linked-list)
* `B` [Doubly Linked List](./data-structures/doubly-linked-list)
* `B` [Queue](./data-structures/queue)
* `B` [Stack](./data-structures/stack)
* `B` [Hash Table](./data-structures/hash-table)
* `B` [Heap](./data-structures/heap) - max and min heap versions
* `B` [Priority Queue](./data-structures/priority-queue)
* `A` [Trie](./data-structures/trie)
* `A` [Tree](./data-structures/tree)
  * `A` [Binary Search Tree](./data-structures/tree/binary-search-tree)
  * `A` [AVL Tree](./data-structures/tree/avl-tree)
  * `A` [Red-Black Tree](./data-structures/tree/red-black-tree)
  * `A` [Segment Tree](./data-structures/tree/segment-tree) - with min/max/sum range queries examples
  * `A` [Fenwick Tree](./data-structures/tree/fenwick-tree) (Binary Indexed Tree)
* `A` [Graph](.data-structures/graph) (both directed and undirected)
* `A` [Disjoint Set](./data-structures/disjoint-set)
* `A` [Bloom Filter](./data-structures/bloom-filter)

## Algorithms

An algorithm is an unambiguous specification of how to solve a class of problems. It is
a set of rules that precisely define a sequence of operations.

`B` - Beginner, `A` - Advanced

### Algorithms by Topic

* **Math**
  * `B` [Bit Manipulation](./algorithms/math/bits) - set/get/update/clear bits, multiplication/division by two, make negative etc.
  * `B` [Factorial](./algorithms/math/factorial)
  * `B` [Fibonacci Number](./algorithms/math/fibonacci) - classic and closed-form versions
  * `B` [Primality Test](./algorithms/math/primality-test) (trial division method)
  * `B` [Euclidean Algorithm](./algorithms/math/euclidean-algorithm) - calculate the Greatest Common Divisor (GCD)
  * `B` [Least Common Multiple](./algorithms/math/least-common-multiple) (LCM)
  * `B` [Sieve of Eratosthenes](./algorithms/math/sieve-of-eratosthenes) - finding all prime numbers up to any given limit
  * `B` [Is Power of Two](./algorithms/math/is-power-of-two) - check if the number is power of two (naive and bitwise algorithms)
  * `B` [Pascal's Triangle](./algorithms/math/pascal-triangle)
  * `B` [Complex Number](./algorithms/math/complex-number) - complex numbers and basic operations with them
  * `B` [Radian & Degree](./algorithms/math/radian) - radians to degree and backwards conversion
  * `B` [Fast Powering](./algorithms/math/fast-powering)
  * `A` [Integer Partition](./algorithms/math/integer-partition)
  * `A` [Liu Hui π Algorithm](./algorithms/math/liu-hui) - approximate π calculations based on N-gons
  * `A` [Discrete Fourier Transform](./algorithms/math/fourier-transform) - decompose a function of time (a signal) into the frequencies that make it up
* **Sets**
  * `B` [Cartesian Product](./algorithms/sets/cartesian-product) - product of multiple sets
  * `B` [Fisher–Yates Shuffle](./algorithms/sets/fisher-yates) - random permutation of a finite sequence
  * `A` [Power Set](./algorithms/sets/power-set) - all subsets of a set (bitwise and backtracking solutions)
  * `A` [Permutations](./algorithms/sets/permutations) (with and without repetitions)
  * `A` [Combinations](./algorithms/sets/combinations) (with and without repetitions)
  * `A` [Longest Common Subsequence](./algorithms/sets/longest-common-subsequence) (LCS)
  * `A` [Longest Increasing Subsequence](./algorithms/sets/longest-increasing-subsequence)
  * `A` [Shortest Common Supersequence](./algorithms/sets/shortest-common-supersequence) (SCS)
  * `A` [Knapsack Problem](./algorithms/sets/knapsack-problem) - "0/1" and "Unbound" ones
  * `A` [Maximum Subarray](./algorithms/sets/maximum-subarray) - "Brute Force" and "Dynamic Programming" (Kadane's) versions
  * `A` [Combination Sum](./algorithms/sets/combination-sum) - find all combinations that form specific sum
* **Strings**
  * `B` [Hamming Distance](./algorithms/string/hamming-distance) - number of positions at which the symbols are different
  * `A` [Levenshtein Distance](./algorithms/string/levenshtein-distance) - minimum edit distance between two sequences
  * `A` [Knuth–Morris–Pratt Algorithm](./algorithms/string/knuth-morris-pratt) (KMP Algorithm) - substring search (pattern matching)
  * `A` [Z Algorithm](./algorithms/string/z-algorithm) - substring search (pattern matching)
  * `A` [Rabin Karp Algorithm](./algorithms/string/rabin-karp) - substring search
  * `A` [Longest Common Substring](./algorithms/string/longest-common-substring)
  * `A` [Regular Expression Matching](./algorithms/string/regular-expression-matching)
* **Searches**
  * `B` [Linear Search](./algorithms/search/linear-search)
  * `B` [Jump Search](./algorithms/search/jump-search) (or Block Search) - search in sorted array
  * `B` [Binary Search](./algorithms/search/binary-search) - search in sorted array
  * `B` [Interpolation Search](./algorithms/search/interpolation-search) - search in uniformly distributed sorted array
* **Sorting**
  * `B` [Bubble Sort](./algorithms/sorting/bubble-sort)
  * `B` [Selection Sort](./algorithms/sorting/selection-sort)
  * `B` [Insertion Sort](./algorithms/sorting/insertion-sort)
  * `B` [Heap Sort](./algorithms/sorting/heap-sort)
  * `B` [Merge Sort](./algorithms/sorting/merge-sort)
  * `B` [Quicksort](./algorithms/sorting/quick-sort) - in-place and non-in-place implementations
  * `B` [Shellsort](./algorithms/sorting/shell-sort)
  * `B` [Counting Sort](./algorithms/sorting/counting-sort)
  * `B` [Radix Sort](./algorithms/sorting/radix-sort)
* **Linked Lists**
  * `B` [Straight Traversal](./algorithms/linked-list/traversal)
  * `B` [Reverse Traversal](./algorithms/linked-list/reverse-traversal)
* **Trees**
  * `B` [Depth-First Search](./algorithms/tree/depth-first-search) (DFS)
  * `B` [Breadth-First Search](./algorithms/tree/breadth-first-search) (BFS)
* **Graphs**
  * `B` [Depth-First Search](./algorithms/graph/depth-first-search) (DFS)
  * `B` [Breadth-First Search](./algorithms/graph/breadth-first-search) (BFS)
  * `B` [Kruskal’s Algorithm](./algorithms/graph/kruskal) - finding Minimum Spanning Tree (MST) for weighted undirected graph
  * `A` [Dijkstra Algorithm](./algorithms/graph/dijkstra) - finding shortest paths to all graph vertices from single vertex
  * `A` [Bellman-Ford Algorithm](./algorithms/graph/bellman-ford) - finding shortest paths to all graph vertices from single vertex
  * `A` [Floyd-Warshall Algorithm](./algorithms/graph/floyd-warshall) - find shortest paths between all pairs of vertices
  * `A` [Detect Cycle](./algorithms/graph/detect-cycle) - for both directed and undirected graphs (DFS and Disjoint Set based versions)
  * `A` [Prim’s Algorithm](./algorithms/graph/prim) - finding Minimum Spanning Tree (MST) for weighted undirected graph
  * `A` [Topological Sorting](./algorithms/graph/topological-sorting) - DFS method
  * `A` [Articulation Points](./algorithms/graph/articulation-points) - Tarjan's algorithm (DFS based)
  * `A` [Bridges](./algorithms/graph/bridges) - DFS based algorithm
  * `A` [Eulerian Path and Eulerian Circuit](./algorithms/graph/eulerian-path) - Fleury's algorithm - Visit every edge exactly once
  * `A` [Hamiltonian Cycle](./algorithms/graph/hamiltonian-cycle) - Visit every vertex exactly once
  * `A` [Strongly Connected Components](./algorithms/graph/strongly-connected-components) - Kosaraju's algorithm
  * `A` [Travelling Salesman Problem](./algorithms/graph/travelling-salesman) - shortest possible route that visits each city and returns to the origin city
* **Cryptography**
  * `B` [Polynomial Hash](./algorithms/cryptography/polynomial-hash) - rolling hash function based on polynomial
* **Uncategorized**
  * `B` [Tower of Hanoi](./algorithms/uncategorized/hanoi-tower)
  * `B` [Square Matrix Rotation](./algorithms/uncategorized/square-matrix-rotation) - in-place algorithm
  * `B` [Jump Game](./algorithms/uncategorized/jump-game) - backtracking, dynamic programming (top-down + bottom-up) and greedy examples
  * `B` [Unique Paths](./algorithms/uncategorized/unique-paths) - backtracking, dynamic programming and Pascal's Triangle based examples
  * `B` [Rain Terraces](./algorithms/uncategorized/rain-terraces) - trapping rain water problem (dynamic programming and brute force versions)
  * `A` [N-Queens Problem](./algorithms/uncategorized/n-queens)
  * `A` [Knight's Tour](./algorithms/uncategorized/knight-tour)

### Algorithms by Paradigm

An algorithmic paradigm is a generic method or approach which underlies the design of a class
of algorithms. It is an abstraction higher than the notion of an algorithm, just as an
algorithm is an abstraction higher than a computer program.

* **Brute Force** - look at all the possibilities and selects the best solution
  * `B` [Linear Search](./algorithms/search/linear-search)
  * `B` [Rain Terraces](./algorithms/uncategorized/rain-terraces) - trapping rain water problem
  * `A` [Maximum Subarray](./algorithms/sets/maximum-subarray)
  * `A` [Travelling Salesman Problem](./algorithms/graph/travelling-salesman) - shortest possible route that visits each city and returns to the origin city
  * `A` [Discrete Fourier Transform](./algorithms/math/fourier-transform) - decompose a function of time (a signal) into the frequencies that make it up
* **Greedy** - choose the best option at the current time, without any consideration for the future
  * `B` [Jump Game](./algorithms/uncategorized/jump-game)
  * `A` [Unbound Knapsack Problem](./algorithms/sets/knapsack-problem)
  * `A` [Dijkstra Algorithm](./algorithms/graph/dijkstra) - finding shortest path to all graph vertices
  * `A` [Prim’s Algorithm](./algorithms/graph/prim) - finding Minimum Spanning Tree (MST) for weighted undirected graph
  * `A` [Kruskal’s Algorithm](./algorithms/graph/kruskal) - finding Minimum Spanning Tree (MST) for weighted undirected graph
* **Divide and Conquer** - divide the problem into smaller parts and then solve those parts
  * `B` [Binary Search](./algorithms/search/binary-search)
  * `B` [Tower of Hanoi](./algorithms/uncategorized/hanoi-tower)
  * `B` [Pascal's Triangle](./algorithms/math/pascal-triangle)
  * `B` [Euclidean Algorithm](./algorithms/math/euclidean-algorithm) - calculate the Greatest Common Divisor (GCD)
  * `B` [Merge Sort](./algorithms/sorting/merge-sort)
  * `B` [Quicksort](./algorithms/sorting/quick-sort)
  * `B` [Tree Depth-First Search](./algorithms/tree/depth-first-search) (DFS)
  * `B` [Graph Depth-First Search](./algorithms/graph/depth-first-search) (DFS)
  * `B` [Jump Game](./algorithms/uncategorized/jump-game)
  * `B` [Fast Powering](./algorithms/math/fast-powering)
  * `A` [Permutations](./algorithms/sets/permutations) (with and without repetitions)
  * `A` [Combinations](./algorithms/sets/combinations) (with and without repetitions)
* **Dynamic Programming** - build up a solution using previously found sub-solutions
  * `B` [Fibonacci Number](./algorithms/math/fibonacci)
  * `B` [Jump Game](./algorithms/uncategorized/jump-game)
  * `B` [Unique Paths](./algorithms/uncategorized/unique-paths)
  * `B` [Rain Terraces](./algorithms/uncategorized/rain-terraces) - trapping rain water problem
  * `A` [Levenshtein Distance](./algorithms/string/levenshtein-distance) - minimum edit distance between two sequences
  * `A` [Longest Common Subsequence](./algorithms/sets/longest-common-subsequence) (LCS)
  * `A` [Longest Common Substring](./algorithms/string/longest-common-substring)
  * `A` [Longest Increasing Subsequence](./algorithms/sets/longest-increasing-subsequence)
  * `A` [Shortest Common Supersequence](./algorithms/sets/shortest-common-supersequence)
  * `A` [0/1 Knapsack Problem](./algorithms/sets/knapsack-problem)
  * `A` [Integer Partition](./algorithms/math/integer-partition)
  * `A` [Maximum Subarray](./algorithms/sets/maximum-subarray)
  * `A` [Bellman-Ford Algorithm](./algorithms/graph/bellman-ford) - finding shortest path to all graph vertices
  * `A` [Floyd-Warshall Algorithm](./algorithms/graph/floyd-warshall) - find shortest paths between all pairs of vertices
  * `A` [Regular Expression Matching](./algorithms/string/regular-expression-matching)
* **Backtracking** - similarly to brute force, try to generate all possible solutions, but each time you generate next solution you test
if it satisfies all conditions, and only then continue generating subsequent solutions. Otherwise, backtrack, and go on a
different path of finding a solution. Normally the DFS traversal of state-space is being used.
  * `B` [Jump Game](./algorithms/uncategorized/jump-game)
  * `B` [Unique Paths](./algorithms/uncategorized/unique-paths)
  * `B` [Power Set](./algorithms/sets/power-set) - all subsets of a set
  * `A` [Hamiltonian Cycle](./algorithms/graph/hamiltonian-cycle) - Visit every vertex exactly once
  * `A` [N-Queens Problem](./algorithms/uncategorized/n-queens)
  * `A` [Knight's Tour](./algorithms/uncategorized/knight-tour)
  * `A` [Combination Sum](./algorithms/sets/combination-sum) - find all combinations that form specific sum
* **Branch & Bound** - remember the lowest-cost solution found at each stage of the backtracking
search, and use the cost of the lowest-cost solution found so far as a lower bound on the cost of
a least-cost solution to the problem, in order to discard partial solutions with costs larger than the
lowest-cost solution found so far. Normally BFS traversal in combination with DFS traversal of state-space
tree is being used.

## Useful Information

### References

### Big O Notation

*Big O notation* is used to classify algorithms according to how their running time or space requirements grow as the input size grows.
On the chart below you may find most common orders of growth of algorithms specified in Big O notation.

![Big O graphs](./assets/big-o-graph.png)

Source: [Big O Cheat Sheet](http://bigocheatsheet.com/).

### Common Data Structure Operations


### Array Sorting Algorithms
