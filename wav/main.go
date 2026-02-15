package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

// WAVHeader represents the structure of a WAV file header
type WAVHeader struct {
	// RIFFチャンク（12バイト）
	ChunkID   [4]byte // "RIFF" - ファイル形式の識別子
	ChunkSize uint32  // ファイルサイズ - 8バイト（ChunkIDとChunkSize自身を除く）
	Format    [4]byte // "WAVE" - WAVEフォーマットであることを示す

	// fmtチャンク（8 + 16 = 24バイト）
	Subchunk1ID   [4]byte // "fmt " - フォーマットチャンクの識別子（スペース含む）
	Subchunk1Size uint32  // fmtチャンクのデータサイズ（PCMでは16バイト）
	AudioFormat   uint16  // 音声フォーマット（1 = PCM、3 = IEEE float、6 = A-law、7 = μ-law）
	NumChannels   uint16  // チャンネル数（1 = モノラル、2 = ステレオ）
	SampleRate    uint32  // サンプリングレート（Hz単位、例：44100、48000、96000）
	ByteRate      uint32  // 1秒あたりのバイト数 = SampleRate × NumChannels × BitsPerSample÷8
	BlockAlign    uint16  // ブロックアライン = NumChannels × BitsPerSample÷8（1サンプルあたりのバイト数）
	BitsPerSample uint16  // 量子化ビット数（8、16、24、32ビット）

	// dataチャンク（8 + データサイズ）
	Subchunk2ID   [4]byte // "data" - データチャンクの識別子
	Subchunk2Size uint32  // 音声データのサイズ = NumSamples × NumChannels × BitsPerSample÷8
}

func main() {
	// ド(C4)の音を1秒間生成
	sampleRate := uint32(44100)    // CD品質のサンプリングレート
	bitsPerSample := uint16(16)    // 16ビット量子化
	numChannels := uint16(1)       // モノラル
	duration := 1.0                // 1秒間
	frequency := 261.63             // ド(C4)の周波数（Hz）
	amplitude := 0.5                // 振幅（最大値の50%）

	// サンプル数を計算
	numSamples := int(float64(sampleRate) * duration)
	
	// 音声データを生成
	audioData := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		// サイン波を生成
		t := float64(i) / float64(sampleRate)
		sample := amplitude * math.Sin(2*math.Pi*frequency*t)
		// 16ビット整数に変換（-32768 ~ 32767）
		audioData[i] = int16(sample * 32767)
	}
	
	// データサイズを計算
	dataSize := uint32(numSamples * int(numChannels) * int(bitsPerSample/8))

	// WAVヘッダーを作成
	header := WAVHeader{
		// RIFFチャンク
		ChunkID:   [4]byte{'R', 'I', 'F', 'F'},
		ChunkSize: 36 + dataSize, // 残り全体のサイズ = 4(Format) + 24(fmtチャンク) + 8(dataチャンクヘッダー) + 0(データ)
		Format:    [4]byte{'W', 'A', 'V', 'E'},

		// fmtチャンク
		Subchunk1ID:   [4]byte{'f', 'm', 't', ' '}, // 最後にスペースが必要
		Subchunk1Size: 16,                           // PCMフォーマットでは16バイト固定
		AudioFormat:   1,                            // 1 = リニアPCM（非圧縮）
		NumChannels:   numChannels,
		SampleRate:    sampleRate,
		ByteRate:      sampleRate * uint32(numChannels) * uint32(bitsPerSample/8), // 44100 * 1 * 2 = 88200
		BlockAlign:    numChannels * bitsPerSample / 8,                             // 1 * 16 / 8 = 2バイト
		BitsPerSample: bitsPerSample,

		// dataチャンク
		Subchunk2ID:   [4]byte{'d', 'a', 't', 'a'},
		Subchunk2Size: dataSize, // 音声データのサイズ
	}

	// ファイルを作成
	file, err := os.Create("c4_note.wav")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()

	// ヘッダーを書き込む
	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		fmt.Printf("ヘッダー書き込みエラー: %v\n", err)
		return
	}

	// 音声データを書き込む
	err = binary.Write(file, binary.LittleEndian, audioData)
	if err != nil {
		fmt.Printf("音声データ書き込みエラー: %v\n", err)
		return
	}

	fmt.Printf("WAVファイル 'c4_note.wav' を作成しました（ド(C4)の音、1秒間）\n")
	fmt.Printf("\n=== WAVファイルの構造 ===\n")
	fmt.Printf("【RIFFチャンク】12バイト\n")
	fmt.Printf("  ChunkID:   'RIFF' (4バイト) - ファイル形式識別子\n")
	fmt.Printf("  ChunkSize: %d (4バイト) - 残りのファイルサイズ\n", header.ChunkSize)
	fmt.Printf("  Format:    'WAVE' (4バイト) - WAVEフォーマット識別子\n")
	
	fmt.Printf("\n【fmtチャンク】24バイト\n")
	fmt.Printf("  Subchunk1ID:   'fmt ' (4バイト) - フォーマットチャンク識別子\n")
	fmt.Printf("  Subchunk1Size: %d (4バイト) - fmtチャンクのデータサイズ\n", header.Subchunk1Size)
	fmt.Printf("  AudioFormat:   %d (2バイト) - 音声フォーマット（1=PCM）\n", header.AudioFormat)
	fmt.Printf("  NumChannels:   %d (2バイト) - チャンネル数\n", header.NumChannels)
	fmt.Printf("  SampleRate:    %d Hz (4バイト) - サンプリングレート\n", header.SampleRate)
	fmt.Printf("  ByteRate:      %d (4バイト) - 1秒あたりのバイト数\n", header.ByteRate)
	fmt.Printf("  BlockAlign:    %d (2バイト) - ブロックアライン\n", header.BlockAlign)
	fmt.Printf("  BitsPerSample: %d (2バイト) - 量子化ビット数\n", header.BitsPerSample)
	
	fmt.Printf("\n【dataチャンク】8バイト\n")
	fmt.Printf("  Subchunk2ID:   'data' (4バイト) - データチャンク識別子\n")
	fmt.Printf("  Subchunk2Size: %d (4バイト) - 音声データのサイズ\n", header.Subchunk2Size)
	
	fmt.Printf("\n音声データ: %d サンプル（%.2f秒）\n", numSamples, duration)
	fmt.Printf("合計ファイルサイズ: %d バイト\n", 44+dataSize)
}