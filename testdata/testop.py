def testop(n):
  bc = dis.opmap[n]
  if bc < dis.HAVE_ARGUMENT:
   print("%s: no argument" % n)
   return
  if bc in dis.hasconst:
   print("%s: has-const" % n)
  elif bc in dis.hasfree:
   print("%s: has-free" %n)
  elif bc in dis.hasname:
   print("%s: has-name"%n)
  elif bc in dis.hasjrel:
   print("%s: has-jump-rel"%n)
  elif bc in dis.hasjabs:
   print("%s: has-jump-abs"%n)
  elif bc in dis.haslocal:
   print("%s: has-local"%n)
  else:
   print("%s: has-argument"%n)
  return


