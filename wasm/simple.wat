(module
;;   (func $i (import "imports" "imported_func") (param i32))
;;   (func (export "exported_func")
;;     i32.const 42
;;     call $i
;;   )

  (func $add (param $lhs i32) (param $rhs i32) (result i32)
    local.get $lhs
    local.get $rhs
    i32.add)


  (export "add" (func $add))
)