/*=============================================================================
#     FileName: rwstream.go
#         Desc: RWStream struct
#       Author: sunminghong
#        Email: allen.fantasy@gmail.com
#     HomePage: http://weibo.com/5d13
#      Version: 0.0.1
#   LastChange: 2013-05-09 17:53:15
#      History:
=============================================================================*/
package net

import (
//    "encoding/binary"
//    "errors"
    "testing"
    "bytes"
)


func Test_NewRWStream(t *testing.T){
    bys :=[]byte{1,2,3,4,5,6,7,8,9,10}

    b := NewRWStream(bys,BigEndian)
    if b.Endian != BigEndian{
        t.Error("NewRWStream error:BigEndian is error",b.Endian)
    }

    _bs := b.Bytes()
    if !bytes.Equal(bys,_bs) {
        t.Error("func Bytes is error:",_bs,bys)
    }

    return
}

func Test_Init(t *testing.T) {
    bytes :=[]byte{1,2,3,4,5,6,7,8,9,10}

    b := NewRWStream(bytes,BigEndian)
    b.Init()

    if (b.last !=0) || (b.end != 0) || (b.off !=0) {
        t.Error("init() is error:last is wrong(0)",b.last)
    }
}

func Test_RW(t *testing.T) {
    bytes :=[]byte{1,2,3,4,5,6,7,8,9,10}

    b := NewRWStream(bytes,BigEndian)
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.Init()
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)

    h,i,j,k,l,m := 1,16,3232,646426464,7777777,-77777777

    for ii:=0;ii<3;ii++ {
        if ii ==2 {
            b.Reset()
        }
    b.WriteByte(byte(h))
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.WriteUint16(uint16(i))
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.WriteUint32(uint32(j))
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.WriteUint64(uint64(k))
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.WriteUint(uint(l))
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    b.WriteInt(m)
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)

    Log(b.buf)

    s := "abcdefghijk"
    b.WriteString(s)
    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)

    h1,err := b.ReadByte()
    if err != nil || int(h1) != h {
        t.Error("ReadByte() error h1=",h1,h)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    i1,err := b.ReadUint16()
    if err != nil || int(i1) != i {
        t.Error("ReadByte() error h1=",i1,i)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    j1,err := b.ReadUint32()
    if err != nil || j1 != uint32(j) {
        t.Error("ReadByte() error h1=",j1,j,err)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    k1,err := b.ReadUint64()
    if err != nil || int(k1) !=k {
        t.Error("ReadByte() error k1=",k1,k)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    l1,err := b.ReadUint()
    if err != nil || int(l1) !=l {
        t.Error("ReadByte() error k1=",l1,l)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    m1,err := b.ReadInt()
    if err != nil || int(m1) !=m {
        t.Error("ReadByte() error k1=",m1,m)
    }

    Log("b.buf Len(),off,end,last=",b.Len(),b.off,b.end,b.last)
    s1,err := b.ReadString()
    if err != nil || s1 !=s {
        t.Error("ReadByte() error k1=",s1,s)
    }

}

}

