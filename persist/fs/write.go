	"github.com/m3db/m3db/digest"
	"github.com/m3db/m3db/generated/proto/schema"
	infoFdWithDigest           digest.FdWithDigestWriter
	indexFdWithDigest          digest.FdWithDigestWriter
	dataFdWithDigest           digest.FdWithDigestWriter
	digestFdWithDigestContents digest.FdWithDigestContentsWriter
	checkpointFilePath         string
	digestBuf    digest.Buffer
	bufferSize int,
		blockSize:                  blockSize,
		filePathPrefix:             filePathPrefix,
		newFileMode:                options.GetNewFileMode(),
		newDirectoryMode:           options.GetNewDirectoryMode(),
		infoBuffer:                 proto.NewBuffer(nil),
		indexBuffer:                proto.NewBuffer(nil),
		varintBuffer:               proto.NewBuffer(nil),
		infoFdWithDigest:           digest.NewFdWithDigestWriter(bufferSize),
		indexFdWithDigest:          digest.NewFdWithDigestWriter(bufferSize),
		dataFdWithDigest:           digest.NewFdWithDigestWriter(bufferSize),
		digestFdWithDigestContents: digest.NewFdWithDigestContentsWriter(bufferSize),
		digestBuf:                  digest.NewBuffer(),
		idxData:                    make([]byte, idxLen),
	shardDir := shardDirPath(w.filePathPrefix, shard)

	var infoFd, indexFd, dataFd, digestFd *os.File
			filepathFromTime(shardDir, blockStart, infoFileSuffix):   &infoFd,
			filepathFromTime(shardDir, blockStart, indexFileSuffix):  &indexFd,
			filepathFromTime(shardDir, blockStart, dataFileSuffix):   &dataFd,
			filepathFromTime(shardDir, blockStart, digestFileSuffix): &digestFd,

	w.infoFdWithDigest.Reset(infoFd)
	w.indexFdWithDigest.Reset(indexFd)
	w.dataFdWithDigest.Reset(dataFd)
	w.digestFdWithDigestContents.Reset(digestFd)

	written, err := w.dataFdWithDigest.WriteBytes(data)
	if _, err := w.indexFdWithDigest.WriteBytes(w.varintBuffer.Bytes()); err != nil {
	if _, err := w.indexFdWithDigest.WriteBytes(entryBytes); err != nil {
	if _, err := w.infoFdWithDigest.WriteBytes(w.infoBuffer.Bytes()); err != nil {
	if err := w.digestFdWithDigestContents.WriteDigests(
		w.infoFdWithDigest.Digest().Sum32(),
		w.indexFdWithDigest.Digest().Sum32(),
		w.dataFdWithDigest.Digest().Sum32(),
	); err != nil {
		return err
	}

	if err := closeAll(
		w.infoFdWithDigest,
		w.indexFdWithDigest,
		w.dataFdWithDigest,
		w.digestFdWithDigestContents,
	); err != nil {
	if err := w.writeCheckpointFile(); err != nil {
		w.err = err
		return err
	}
	return nil
	defer fd.Close()
	if err := w.digestBuf.WriteDigestToFile(fd, w.digestFdWithDigestContents.Digest().Sum32()); err != nil {
		return err
	}