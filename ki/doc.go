// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package ki provides the base element of GoKi Trees: Ki = Tree in Japanese, and
"Key" in English -- powerful tree structures supporting scenegraphs, programs,
parsing, etc.

The Node struct that implements the Ki interface, which can be used as an
embedded type (or a struct field) in other structs to provide core tree
functionality, including:

	* Parent / Child Tree structure -- each Node can ONLY have one parent.
      Node struct's can also have Node fields -- these are functionally like
      fixed auto-named children.

	* Paths for locating Nodes within the hierarchy -- key for many use-cases,
      including ability to convert pointers to/from strings for IO and robust
      deep copy and move functions.  The path separator is / for children and
      . for fields.

	* Apply a function across nodes up or down a tree (natural "me first",
      breadth-first, depth-first) -- very flexible for tree walking.

	* Generalized I/O -- can Save and Load the Tree as JSON, XML, etc --
      including pointers which are saved using paths and automatically
      cached-out after loading -- enums also bidirectionally convertable to
      strings using enum type registry in kit package.

	* Robust deep copy, clone, move of nodes, with automatic pointer updating.

	* Signal sending and receiving between Nodes (simlar to Qt Signals /
      Slots) -- setup connections once and then emit signals to all receivers
      when relevant event happens.

	* Robust state updating -- wrap updates in UpdateStart / End, and signals
      are blocked until the final end, at the highest affected level in the
      tree, at which point a single update signal is sent -- automatically
      gives the minimal update.

	* Properties (as a string-keyed map) with property inheritance, including
      type-level properties via kit type registry.

*/
package ki
