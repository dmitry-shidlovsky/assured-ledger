// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package cryptkit

import (
	"bytes"
	"hash/crc32"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/insolar/assured-ledger/ledger-core/vanilla/longbits"
)

func TestCompareNil(t *testing.T) {
	var s1, s2 Signature
	require.Equal(t, 0, s1.FixedByteSize())
	require.True(t, s1.Equals(s2))
}

func TestCompareNilZero(t *testing.T) {
	require.Equal(t, 0, Digest{}.FixedByteSize())
	require.Equal(t, DigestMethod(""), Digest{}.GetDigestMethod())
	require.Nil(t, Digest{}.AsDigestHolder())
	require.True(t, Digest{}.IsEmpty())
	require.True(t, Digest{}.IsZero())

	require.Equal(t, 0, Signature{}.FixedByteSize())
	require.Equal(t, SignatureMethod(""), Signature{}.GetSignatureMethod())
	require.Nil(t, Signature{}.AsSignatureHolder())
	require.True(t, Signature{}.IsEmpty())

	require.Equal(t, 0, SignatureKey{}.FixedByteSize())
	require.Equal(t, SignatureMethod(""), SignatureKey{}.GetSignatureKeyMethod())
	require.Equal(t, SigningMethod(""), SignatureKey{}.GetSigningMethod())
	require.True(t, SignatureKey{}.IsEmpty())

	require.Equal(t, 0, SignedData{}.FixedByteSize())
	require.True(t, SignedData{}.IsEmpty())
}

func TestIsSymmetric(t *testing.T) {
	require.True(t, SymmetricKey.IsSymmetric())

	require.False(t, PublicAsymmetricKey.IsSymmetric())
}

func TestIsSecret(t *testing.T) {
	require.True(t, SymmetricKey.IsSecret())

	require.False(t, PublicAsymmetricKey.IsSecret())
}

func TestSignedBy(t *testing.T) {
	td := "testDigest"
	ts := "testSign"
	require.Equal(t, SignatureMethod(strings.Join([]string{td, ts}, "/")), DigestMethod(td).SignedBy(SigningMethod(ts)))
}

func TestDigestMethodString(t *testing.T) {
	td := "test"
	d := DigestMethod(td)
	require.Equal(t, td, d.String())
}

func TestSignMethodString(t *testing.T) {
	ts := "test"
	s := SigningMethod(ts)
	require.Equal(t, ts, s.String())
}

func TestDigestMethod(t *testing.T) {
	td := "testDigest"
	ts := "testSign"
	sep := "/"
	require.Equal(t, DigestMethod(td), SignatureMethod(strings.Join([]string{td, ts}, sep)).DigestMethod())

	emptyDigMethod := DigestMethod("")
	require.Equal(t, emptyDigMethod, SignatureMethod("testSignature").DigestMethod())

	require.Equal(t, emptyDigMethod, SignatureMethod(strings.Join([]string{td, ts, "test"}, sep)).DigestMethod())
}

func TestSignMethod(t *testing.T) {
	td := "testDigest"
	ts := "testSign"
	sep := "/"
	require.Equal(t, SigningMethod(ts), SignatureMethod(strings.Join([]string{td, ts}, sep)).SignMethod())

	emptySignMethod := SigningMethod("")
	require.Equal(t, emptySignMethod, SignatureMethod("testSignature").SignMethod())

	require.Equal(t, emptySignMethod, SignatureMethod(strings.Join([]string{td, ts, "test"}, sep)).SignMethod())
}

func TestString(t *testing.T) {
	ts := "test"
	sm := SignatureMethod(ts)
	require.Equal(t, ts, sm.String())
}

func TestCopyOfDigest(t *testing.T) {
	d := &Digest{digestMethod: "test"}
	fd := longbits.NewFoldableReaderMock(t)
	fd.FixedByteSizeMock.Set(func() int { return 0 })
	fd.CopyToMock.Set(func(p []byte) int { return 0 })
	d.hFoldReader = fd
	cd := d.CopyOfDigest()
	require.Equal(t, cd.digestMethod, d.digestMethod)
}

func TestDigestEquals(t *testing.T) {
	bits := longbits.NewBits64(0)
	d := NewDigest(&bits, "")
	dc := NewDigest(&bits, "")
	require.True(t, d.Equals(dc.AsDigestHolder()))
}

func TestNewDigest(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	method := DigestMethod("test")
	d := NewDigest(fd, method)
	require.Equal(t, fd, d.hFoldReader)

	require.Equal(t, method, d.digestMethod)
}

func TestSignWith(t *testing.T) {
	ds := NewDigestSignerMock(t)
	sm := SignatureMethod("test")
	ds.SignDigestMock.Set(func(Digest) Signature { return Signature{signatureMethod: sm} })
	d := &Digest{}
	sd := d.SignWith(ds)
	require.Equal(t, sm, sd.GetSignatureMethod())
}

func TestDigestString(t *testing.T) {
	require.True(t, Digest{}.String() != "")
}

func TestCopyOfSignature(t *testing.T) {
	s := &Signature{signatureMethod: "test"}
	fd := longbits.NewFoldableReaderMock(t)
	fd.FixedByteSizeMock.Set(func() int { return 0 })
	fd.CopyToMock.Set(func(p []byte) int { return 0 })
	s.hFoldReader = fd
	cs := s.CopyOfSignature()
	require.Equal(t, cs.signatureMethod, s.signatureMethod)
}

func TestNewSignature(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	method := SignatureMethod("test")
	s := NewSignature(fd, method)
	require.Equal(t, fd, s.hFoldReader)

	require.Equal(t, method, s.signatureMethod)
}

func TestSignatureEquals(t *testing.T) {
	bits := longbits.NewBits64(0)
	s := NewSignature(&bits, "")
	sc := NewSignature(&bits, "")
	require.True(t, s.Equals(sc.AsSignatureHolder()))
}

func TestSignGetSignatureMethod(t *testing.T) {
	ts := SignatureMethod("test")
	signature := NewSignature(nil, ts)
	require.Equal(t, ts, signature.GetSignatureMethod())
}

func TestAsSignatureHolder(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	s := Signature{hFoldReader: fd, signatureMethod: "test"}
	sh := s.AsSignatureHolder()
	require.Equal(t, sh.GetSignatureMethod(), s.signatureMethod)

	require.Implements(t, (*SignatureHolder)(nil), sh)
}

func TestSignatureString(t *testing.T) {
	require.True(t, Signature{}.String() != "")
}

func TestNewSignedDigest(t *testing.T) {
	d := Digest{digestMethod: "testDigest"}
	s := Signature{signatureMethod: "testSignature"}
	sd := NewSignedDigest(d, s)
	require.Equal(t, d.digestMethod, sd.digest.digestMethod)

	require.Equal(t, s.signatureMethod, sd.GetSignatureMethod())
}

func TestCopyOfSignedDigest(t *testing.T) {
	d := Digest{digestMethod: "testDigest"}
	fd1 := longbits.NewFoldableReaderMock(t)
	fd1.FixedByteSizeMock.Set(func() int { return 0 })
	fd1.CopyToMock.Set(func(p []byte) int { return 0 })
	d.hFoldReader = fd1

	s := Signature{signatureMethod: "testSignature"}
	fd2 := longbits.NewFoldableReaderMock(t)
	fd2.FixedByteSizeMock.Set(func() int { return 0 })
	fd2.CopyToMock.Set(func(p []byte) int { return 0 })
	s.hFoldReader = fd2
	sd := NewSignedDigest(d, s)
	sdc := sd.CopyOfSignedDigest()
	require.Equal(t, sdc.digest.digestMethod, sd.digest.digestMethod)

	require.Equal(t, sdc.GetSignatureMethod(), sd.GetSignatureMethod())
}

func TestSignedDigestEquals(t *testing.T) {
	dBits := longbits.NewBits64(0)
	d := NewDigest(&dBits, "")

	sBits1 := longbits.NewBits64(0)
	s := NewSignature(&sBits1, "")

	sd1 := NewSignedDigest(d, s)
	sd2 := NewSignedDigest(d, s)
	require.True(t, sd1.Equals(&sd2))

	sBits2 := longbits.NewBits64(1)
	sd2 = NewSignedDigest(d, NewSignature(&sBits2, ""))
	require.False(t, sd1.Equals(&sd2))
}

func TestGetDigest(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	d := Digest{hFoldReader: fd, digestMethod: "test"}
	s := Signature{}
	sd := NewSignedDigest(d, s)
	require.Equal(t, fd, sd.GetDigest().hFoldReader)

	require.Equal(t, d.digestMethod, sd.GetDigest().digestMethod)
}

func TestGetSignature(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	d := Digest{}
	s := Signature{hFoldReader: fd, signatureMethod: "test"}
	sd := NewSignedDigest(d, s)
	require.Equal(t, fd, sd.GetSignature().hFoldReader)

	require.Equal(t, s.signatureMethod, sd.GetSignature().signatureMethod)
}

func TestGetDigestHolder(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	d := Digest{hFoldReader: fd, digestMethod: "testDigest"}
	s := Signature{hFoldReader: fd, signatureMethod: "testSignature"}
	sd := NewSignedDigest(d, s)
	require.Equal(t, d.AsDigestHolder(), sd.GetDigestHolder())
}

func TestGetSignatureHolder(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	d := Digest{hFoldReader: fd, digestMethod: "testDigest"}
	s := Signature{hFoldReader: fd, signatureMethod: "testSignature"}
	sd := NewSignedDigest(d, s)
	require.Equal(t, s.AsSignatureHolder(), sd.GetSignatureHolder())
}

func TestSignedDigGetSignatureMethod(t *testing.T) {
	s := Signature{signatureMethod: "test"}
	sd := NewSignedDigest(Digest{}, s)
	require.Equal(t, s.signatureMethod, sd.GetSignatureMethod())
}

func TestIsVerifiableBy(t *testing.T) {
	sd := NewSignedDigest(Digest{}, Signature{})
	sv := NewSignatureVerifierMock(t)
	supported := false
	sv.IsSignOfSignatureMethodSupportedMock.Set(func(SignatureMethod) bool { return *(&supported) })
	require.False(t, sd.IsVerifiableBy(sv))

	supported = true
	require.True(t, sd.IsVerifiableBy(sv))
}

func TestVerifyWith(t *testing.T) {
	sd := NewSignedDigest(Digest{}, Signature{})
	sv := NewSignatureVerifierMock(t)
	valid := false
	sv.IsValidDigestSignatureMock.Set(func(DigestHolder, SignatureHolder) bool { return *(&valid) })
	require.False(t, sd.VerifyWith(sv))

	valid = true
	require.True(t, sd.VerifyWith(sv))
}

func TestSignedDigestString(t *testing.T) {
	require.True(t, NewSignedDigest(Digest{}, Signature{}).String() != "")
}

func TestAsSignedDigestHolder(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	d := Digest{hFoldReader: fd, digestMethod: "testDigest"}
	s := Signature{hFoldReader: fd, signatureMethod: "testSignature"}
	sd := NewSignedDigest(d, s)
	sdh := sd.AsSignedDigestHolder()

	require.Equal(t, sdh.GetSignatureMethod(), s.signatureMethod)

	require.Implements(t, (*SignedDigestHolder)(nil), sdh)
}

func TestNewSignedData(t *testing.T) {
	bits := longbits.NewBits64(0)
	d := Digest{digestMethod: "testDigest"}
	s := Signature{signatureMethod: "testSignature"}
	sd := NewSignedData(&bits, d, s)
	require.Equal(t, &bits, sd.hWriterTo)

	require.Equal(t, d, sd.hSignedDigest.digest)

	require.Equal(t, s, sd.hSignedDigest.signature)
}

func TestGetSignedDigest(t *testing.T) {
	bits := longbits.NewBits64(0)
	d := Digest{digestMethod: "testDigest"}
	s := Signature{signatureMethod: "testSignature"}
	sd := NewSignedData(&bits, d, s)
	signDig := sd.GetSignedDigest()
	require.Equal(t, d, signDig.digest)

	require.Equal(t, s, signDig.signature)
}

func TestWriteTo(t *testing.T) {
	bits1 := longbits.NewBits64(1)
	d := Digest{digestMethod: "testDigest"}
	s := Signature{signatureMethod: "testSignature"}
	sd := NewSignedData(&bits1, d, s)
	buf := &bytes.Buffer{}
	n, err := sd.WriteTo(buf)
	require.Equal(t, int64(8), n)

	require.Nil(t, err)
}

func TestSignedDataString(t *testing.T) {
	bits := longbits.NewBits64(0)
	require.NotEmpty(t, NewSignedData(&bits, Digest{}, Signature{}).String())
}

func TestNewSignatureKey(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	ts := SignatureMethod("testSign")
	kt := PublicAsymmetricKey
	sk := NewSignatureKey(fd, ts, kt)
	require.Equal(t, fd, sk.hFoldReader)

	require.Equal(t, ts, sk.signatureMethod)

	require.Equal(t, kt, sk.keyType)
}

func TestGetSignMethod(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	ts := "testSign"
	sk := NewSignatureKey(fd, SignatureMethod(strings.Join([]string{"testDigest", ts}, "/")), PublicAsymmetricKey)
	require.Equal(t, ts, sk.GetSigningMethod().String())
}

func TestGetSignatureKeyMethod(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	s := strings.Join([]string{"testDigest", "testSign"}, "/")
	sk := NewSignatureKey(fd, SignatureMethod(s), PublicAsymmetricKey)
	require.Equal(t, s, sk.GetSignatureKeyMethod().String())
}

func TestGetSignatureKeyType(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	kt := PublicAsymmetricKey
	sk := NewSignatureKey(fd, "test", kt)
	require.Equal(t, kt, sk.GetSignatureKeyType())
}

func TestEquals(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	fd.FixedByteSizeMock.Set(func() int { return 0 })
	fd.WriteToMock.Set(func(io.Writer) (int64, error) { return 0, nil })
	sk1 := NewSignatureKey(fd, "test", PublicAsymmetricKey)
	sk2 := NewSignatureKey(fd, "test", PublicAsymmetricKey)
	require.True(t, sk1.Equals(&sk2))
}

func TestSignatureKeyString(t *testing.T) {
	fd := longbits.NewFoldableReaderMock(t)
	sk := NewSignatureKey(fd, "test", PublicAsymmetricKey)
	require.NotEmpty(t, sk.String())
}

func TestSignDataByDataSigner(t *testing.T) {
	signer := NewDataSignerMock(t)
	signer.GetDigestSizeMock.Return(4)
	signer.GetDigestMethodMock.Return("testMethod")
	signer.NewHasherMock.Return(DigestHasher{BasicDigester: signer, Hash: crc32.NewIEEE()})
	data := longbits.WrapStr("testString")
	hasher := crc32.NewIEEE()
	_, _ = data.WriteTo(hasher)

	signer.SignDigestMock.Expect(NewDigest(longbits.NewMutableFixedSize(hasher.Sum(nil)), "testMethod"))
	signer.SignDigestMock.Return(Signature{})
	sdata := SignDataByDataSigner(data, signer)
	require.True(t, longbits.Equal(data, sdata))
}

func TestNewZeroSizeDigest(t *testing.T) {
	d := NewZeroSizeDigest("test")
	require.Equal(t, 0, d.FixedByteSize())
	require.Equal(t, DigestMethod("test"), d.GetDigestMethod())
	require.NotNil(t, d.AsDigestHolder())
	require.False(t, d.IsEmpty())
	require.False(t, d.IsZero())
}